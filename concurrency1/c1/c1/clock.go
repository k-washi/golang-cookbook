package c1

import (
	"io"
	"net"
	"os"
	"time"

	"github.com/k-washi/bss-utils/logger"
)

func StartServer() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			logger.Log.Fatal(err.Error())
			continue
		}
		handleConn(conn)

	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func StartClient() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		logger.Log.Fatal(err.Error())
	}
}
