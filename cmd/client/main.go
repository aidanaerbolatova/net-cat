package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8989")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		mustcopy(os.Stdout, conn)
		done <- struct{}{}
	}()

	mustcopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustcopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
