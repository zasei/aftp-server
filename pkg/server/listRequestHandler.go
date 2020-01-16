package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func handleListRequest(request Request, conn net.Conn) {
	createResponse := Response{
		protocol:   ProtocolVersion,
		statusCode: OK,
		headers:    nil,
		message:    listDirectory(request.parameters[0]),
	}

	fmt.Printf("handleLisRequest with createResponse: %s\n", createResponse)
	doHandle(createResponse, conn)
}

func listDirectory(path string) string {
	//files, err := ioutil.ReadDir("./" + FileDir + path)
	files, err := ioutil.ReadDir(FileDir + path)

	if err != nil {
		log.Fatal(err)
	}

	var results strings.Builder

	for _, f := range files {
		results.WriteString(f.Name() + "\n")
	}

	return results.String()
}
