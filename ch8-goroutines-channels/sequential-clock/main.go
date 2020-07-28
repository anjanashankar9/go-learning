package main

import (
	"io"
	"log"
	"net"
	"time"
)

/*
Networking is a natural domain in which to use concurrency since
servers typically handle many connections from their clients at once,
each client being essentially independent of the others.
*/

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err) //e.g. Connection aborted
			continue
		}
		go handleConn(conn) //handle one connection at a time
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
