/**
 * datetime: 2017-05-30
 * netcat - go_clock的客户端
 * The Go Programming Language example
 */
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//server: go run go_clock.go
//client: go run go_netcat.go

//
// Client: :  127.0.0.1:62367
// Client: :  127.0.0.1:62563
// Client: :  127.0.0.1:62585
//
