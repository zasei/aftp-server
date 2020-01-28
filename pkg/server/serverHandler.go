package server

import (
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"net"
	"strings"
)

// TODO: Change to other relative path
const FileDir = "./files"

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
		handleBadRequest(conn, "invalid use of protocol")
	}

	parsedRequest := dom.ParseRequest(string(buf))

	parsedRequest.PrintRequest()

	if !strings.Contains(parsedRequest.Protocol, dom.ProtocolVersion) {
		handleBadRequest(conn, fmt.Sprintf("unknown protocol %s, wanted %s", parsedRequest.Protocol, dom.ProtocolVersion))
	}

	switch parsedRequest.Method {
	case dom.GET:
		handleGetRequest(parsedRequest, conn)
	case dom.LIST:
		handleListRequest(parsedRequest, conn)
	case dom.PUT:
		handlePutRequest(parsedRequest, conn)
	case dom.DELETE:
		handleDeleteRequest(parsedRequest, conn)
	default:
		handleBadRequest(conn, "unknown/unsupported method: "+parsedRequest.Method)
	}
}

func doHandle(response dom.Response, conn net.Conn) {
	fmt.Println("doHandle")
	createdResponse := response.CreateResponse()
	_, _ = conn.Write([]byte(createdResponse))
	err := conn.Close()
	if err != nil {
		fmt.Printf("Error sending response %s\n", err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
