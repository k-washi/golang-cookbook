package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	ChatOneLineSize = 512
	ChatLogLength   = 100
	SendBufferSize  = 4 * 1024 * 1024
)

/*
+ 特定のクライアントへの送信に固執せず、他のクライアントへの送信が送れないようにする
+ 切断のタイミングは、 1.ルーム側からのkick, 2. 受信エラー, 3. 送信エラー
+ 送信待ちのメッセージを無限にバッファリングすると、メッセージを送信するだけで受信しないようなクライアントがいる場合にメモリを使い切ってしまう場合がある。結果、用意にDoS攻撃が可能になっていしまう。

バッファリング: 一時的にバッファーメモリーなどにデータを蓄え、データ処理速度や処理にかかる時間の違いを調整すること
[D]oS攻撃](https://cybersecurity-jp.com/security-measures/18262): 攻撃目標であるサイトやサーバに対して大量のデータを送り付けることで行われる攻撃
OOM Killer: メモリが不足してシステムが停止する恐れがある際、メモリリソースを多く消費しているプロセスを強制的に殺します
*/

type Room struct {
	Join   chan *Client // new clients
	Closed chan *Client // leaved clients
	Recv   chan string  // received message from clients.
	Purge  chan bool    // kick all clients in this room.
	Stop   chan bool    // close this room.
	// slice の resize コストが接続数に比例して増えるのを防ぐため map を使う.
	clients map[*Client]bool
	log     []string
}

func newRoom() *Room {
	r := &Room{
		Join:    make(chan *Client),
		Closed:  make(chan *Client),
		Recv:    make(chan string),
		Purge:   make(chan bool),
		Stop:    make(chan bool),
		clients: make(map[*Client]bool),
	}
	go r.run()
	return r
}

func (r *Room) run() {
	defer log.Println("Room closed")
	for {
		select {
		case c := <-r.Join:
			log.Printf("Room: %v is joined", c)
			if err := r.sendLog(c); err != nil {
				log.Println(err)
				c.Stop()
			} else {
				r.clients[c] = true
			}

		case c := <-r.Closed:
			log.Printf("Room: %v has been closed", c)
			delete(r.clients, c)

		case msg := <-r.Recv:
			log.Printf("Room: Received %#v", msg) //%#vはgoの構文のまま表示
			r.appendLog(msg)
			for c := range r.clients {
				if err := c.Send(msg); err != nil {
					log.Println(err)
					c.Stop()
					delete(r.clients, c)
				}
			}
		case <-r.Purge:
			log.Printf("Purge all clients")
			r.purge()
		case <-r.Stop:
			log.Println("Closeing room...")
			r.purge()
			return

		}
	}
}

func (r *Room) appendLog(msg string) {
	r.log = append(r.log, msg)
	if len(r.log) > ChatLogLength {
		r.log = r.log[len(r.log)-ChatLogLength:]
	}
}

func (r *Room) sendLog(c *Client) error {
	for _, msg := range r.log {
		if err := c.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func (r *Room) purge() {
	for c := range r.clients {
		c.Stop()
		delete(r.clients, c)
	}
}

type Client struct {
	m       sync.Mutex
	id      int
	recv    chan string
	closed  chan *Client
	conn    *net.TCPConn
	stop    bool
	cond    sync.Cond
	sendBuf *bytes.Buffer
}

func (c *Client) String() string {
	return fmt.Sprintf("Client(%v)", c.id)
}

var lastClientId = 0
var clientWait sync.WaitGroup

func newClient(r *Room, conn *net.TCPConn) *Client {
	lastClientId++
	cl := &Client{
		m:       sync.Mutex{},
		id:      lastClientId,
		recv:    r.Recv,
		closed:  r.Closed,
		conn:    conn,
		stop:    false,
		sendBuf: &bytes.Buffer{},
	}

	cl.cond.L = &cl.m
	clientWait.Add(1)
	go cl.sender()
	go cl.receiver()
	log.Printf("%v is created", cl)
	return cl

}

func (c *Client) Send(msg string) error {
	c.m.Lock()
	defer c.m.Unlock()

	if c.sendBuf.Len()+len(msg) > SendBufferSize {
		return errors.New("Buffer full")
	}
	c.sendBuf.WriteString(msg)
	c.cond.Signal()
	return nil
}

func (c *Client) Stop() {
	c.m.Lock()
	c.stop = true
	c.m.Unlock()
	c.cond.Signal()
}

func (c *Client) sender() {
	defer func() {
		if err := c.conn.Close(); err != nil {
			log.Println(err)
		}
		log.Printf("%v is closed", c)
		clientWait.Done()
		c.closed <- c
	}()

	buf := &bytes.Buffer{}

	c.m.Lock()
	for {
		if c.stop {
			return
		}
		if c.sendBuf.Len() == 0 {
			c.cond.Wait()
			continue
		}

		buf, c.sendBuf = c.sendBuf, buf
		c.m.Unlock()

		_, err := c.conn.Write(buf.Bytes())
		if err != nil {
			log.Println(c, err)
			return
		}
		buf.Reset()
		c.m.Lock()
	}
}

func (c *Client) receiver() {
	defer c.Stop()

	reader := bufio.NewReaderSize(c.conn, ChatOneLineSize)
	for {
		msg, err := reader.ReadString(byte('\n'))
		if msg != "" {
			c.recv <- msg
		}
		if err != nil {
			log.Println("receiver: ", err)
			return
		}
	}
}

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	log.Println("PID: ", os.Getegid())

	room := newRoom()

	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:5056")
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			//Acceptで他のリクエストを受付
			conn, err := l.AcceptTCP()
			if err != nil {
				log.Println(err)
				l.Close()
				return
			}
			room.Join <- newClient(room, conn)
		}
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGUSR1, syscall.SIGTERM, os.Interrupt)
	for sig := range sigc {
		switch sig {
		case syscall.SIGUSR1:
			room.Purge <- true
		case syscall.SIGTERM, os.Interrupt:
			l.Close()
			room.Stop <- true
			clientWait.Wait()
			return
		}
	}
}
