package server

import (
	dom "aftp-server/pkg/domain"
	"fmt"
	"net"
	"strings"
)

// TODO: Change to other relative path
const FileDir = "/home/rcomanne/go/src/aftp-server/files"

func HandleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	//Read the incoming connection into the buffer
	_, err := conn.Read(buf)

	requestString := strings.Fields(string(buf))

	if err != nil {
		fmt.Println("Error reading", err.Error())
		handleServerError(conn)
	}

	if len(requestString) < 3 {
		handleBadRequest(conn)
	}
	parsedRequest := dom.ParseRequest(requestString)
	switch parsedRequest.Method {
	case dom.GET:
		handleGetRequest(parsedRequest, conn)
	case dom.LIST:
		handleListRequest(parsedRequest, conn)
	default:
		handleBadRequest(conn)
	}
}

func doHandle(response dom.Response, conn net.Conn) {
	fmt.Println("doHandle")
	createdResponse := response.CreateResponse()
	_, _ = conn.Write([]byte(createdResponse))
	err := conn.Close()
	if err != nil {
		fmt.Printf("Error sending response %s", err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
