package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"net-cat-v0.1/server"
)

func main() {
	host := "0.0.0.0"
	port := 8989

	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 2 {

		_, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalln("Invalid port number")
		}
		port, _ = strconv.Atoi(os.Args[1])
	}

	server.StartServer(host, port)
}
