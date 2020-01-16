package main

import (
	"aftp-server/pkg/server"
	"fmt"
	"net"
	"os"
)

// server settings
const (
	HOST = "127.0.0.1"
	PORT = "1337"
	TYPE = "tcp"
)

func main() {

	l, err := net.Listen(TYPE, HOST+":"+PORT)

	if err != nil {
		fmt.Printf("Error listening on %s:%s\n", HOST, PORT)
		os.Exit(1)
	}

	defer l.Close()

	fmt.Printf("Listening on %s:%s\n", HOST, PORT)

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting new connection", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine
		go server.HandleRequest(conn)
	}

}
