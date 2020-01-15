package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func HandleRequest(dirs []string) {
	request := Request{
		method:    LIST,
		protocol:  VERSION,
		headers:   nil,
		parameter: dirs[0],
	}

	doHandle(request)
}

func doHandle(request Request) {
	// Set up connection to remote
	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Printf("Error while connecting %s\n", err)
		os.Exit(1)
	}

	// Create the request as a String literal
	requestString := createRequest(request)
	fmt.Printf("Created request: \n%s\n", requestString)

	// Write to remote connection
	code, err := fmt.Fprintf(conn, requestString)
	if err != nil {
		fmt.Printf("Error writing to remote %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Int? %d\n", code)

	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}

type Request struct {
	method    string
	protocol  string
	headers   []string
	parameter string
}

type Response struct {
	protocol   string
	statusCode string
	headers    []string
	message    string
}
