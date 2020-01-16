package main

import (
	dom "aftp-server/pkg/domain"
	"aftp-server/pkg/server"
	"fmt"
	"net"
	"os"
)

func main() {

	l, err := net.Listen(dom.TYPE, dom.HOST+":"+dom.PORT)

	if err != nil {
		fmt.Printf("Error listening on %s:%s\n", dom.HOST, dom.PORT)
		os.Exit(1)
	}

	defer l.Close()

	fmt.Printf("Listening on %s:%s\n", dom.HOST, dom.PORT)

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
