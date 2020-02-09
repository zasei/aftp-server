package server

import (
	"encoding/gob"
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"log"
	"net"
	"strings"
)

// TODO: Change to other relative path
//const FileDir = "/opt/aftp-server"
const FileDir = "./files"

func HandleRequest(conn net.Conn) {

	dec := gob.NewDecoder(conn)

	var parsedRequest dom.Request

	err := dec.Decode(&parsedRequest)

	if err != nil {
		log.Fatal("decode error:", err)
	}

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
