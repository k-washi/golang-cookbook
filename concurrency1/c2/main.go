package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
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
			}

		case c := <-r.Closed:
			log.Printf("Room: %v has been closed", c)

		case msg := <-r.Recv:
			log.Printf("Room: Received %#v", msg) //%#vはgoの構文のまま表示
		case <-r.Purge:
			log.Printf("Purge all clients")
		case <-r.Stop:
			log.Println("Closeing room...")

		}
	}
}

func (r *Room) sendLog(c *Client) error {
	for _, msg := range r.log {
		//client send msg
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

func newClient(r *Room, conn *net.TCPConn)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	log.Println("PID: ", os.Getegid())
}
