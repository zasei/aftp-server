package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func handleListRequest(request Request, conn net.Conn) {
	response := Response{
		protocol:   VERSION,
		statusCode: OK,
		headers:    nil,
		message:    listDirectory(request.parameters[0]),
	}

	fmt.Print(response)
	doHandle(response, conn)
}

func listDirectory(path string) string {
	files, err := ioutil.ReadDir("./" + FILE_DIR + path)

	if err != nil {
		log.Fatal(err)
	}

	var results strings.Builder

	for _, f := range files {
		results.WriteString(f.Name() + "\n")
	}

	return results.String()
}
