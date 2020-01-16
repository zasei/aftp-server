package client

import (
	dom "aftp-server/pkg/domain"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

// Handle the actual network connection and parsing
func doHandle(request dom.Request) {
	fmt.Println("doHandle")
	// Set up connection to remote
	conn, err := net.Dial(dom.TYPE, dom.HOST+":"+dom.PORT)
	if err != nil {
		fmt.Printf("Error while connecting %s\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	// Create the request as a String literal
	requestString := request.RequestToString()

	// Write to remote connection
	fmt.Fprintf(conn, requestString)

	// create buffer and copy bytes to it
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	receivedResponse := dom.ParseResponse(buf)

	receivedResponse.PrintResponse()
}
