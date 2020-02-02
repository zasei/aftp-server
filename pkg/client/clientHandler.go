package client

import (
	"bytes"
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"io"
	"net"
	"os"
)

// Handle the actual network connection and parsing
func doHandle(request dom.Request) dom.Response {
	// Set up connection to remote
	conn, err := net.Dial(dom.TYPE, dom.HOST+":"+dom.PORT)
	if err != nil {
		fmt.Printf("Error while connecting %s\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	// Create the request as a String literal
	requestString := request.RequestToString()

	// Uncomment this line to view the request string
	//fmt.Printf("Request String: %s\n", requestString)

	// Write to remote connection
	fmt.Fprintf(conn, requestString)

	// create buffer and copy bytes to it
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	receivedResponse := dom.ParseResponse(buf)

	// uncomment this line to show the full parse response
	//receivedResponse.PrintResponse()

	return receivedResponse
}
