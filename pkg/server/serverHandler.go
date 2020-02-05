package server

import (
	"bufio"
	"bytes"
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"net"
	"strings"
)

// TODO: Change to other relative path
//const FileDir = "/opt/aftp-server"
const FileDir = "./files"

func HandleRequest(conn net.Conn) {

	// Create a buffered reader for the connection
	connbuf := bufio.NewReader(conn)
	// Read first byte and set buffer
	b, err := connbuf.ReadByte()

	if err != nil {
		fmt.Println("Unable to read from connection.")
	}

	if connbuf.Buffered() > 0 {
		var requestData []byte
		requestData = append(requestData, b)
		for connbuf.Buffered() > 0 {
			// Read byte by byte until empty or error
			b, err := connbuf.ReadByte()
			if err == nil {
				requestData = append(requestData, b)
			} else {
				fmt.Println("... Unreadable character ...")
				fmt.Println(b)
			}
		}
		fullRequest := string(requestData)
		fmt.Println(fullRequest)

		requestString := strings.Fields(string(requestData))

		if len(requestString) < 3 {
			handleBadRequest(conn, "invalid use of protocol")
		}

		parsedRequest := dom.ParseRequest(string(bytes.Trim(requestData, "\x00")))

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
