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

/*
The program reads data from the connection and writes it to
the standard output until an end of file condition or an
error occurs.

Running two instances shows that the second client must wait
until the first client is finished because the server is
sequential, it deals with only one client at a time.

Just one change is needed to make the server concurrent: adding the go
keyword to the call to handleConn causes each call to run in
its own goroutine.
*/
