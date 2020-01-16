package server

import (
	dom "aftp-server/pkg/domain"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func handleListRequest(request dom.Request, conn net.Conn) {
	createResponse := dom.Response{
		Protocol:   dom.ProtocolVersion,
		StatusCode: dom.OK,
		Headers:    nil,
		Message:    listDirectory(request.Parameter),
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
