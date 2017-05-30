/**
 * datetime: 2017-05-30
 * 并发的clock 服务
 * The Go Programming Language example
 */
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	//创建一个对象监听网络连接,协议tcp, 端口8000
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//接收新的连接，如果没有则阻塞
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
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
		time.Sleep(time.Second * 1)
	}
}
