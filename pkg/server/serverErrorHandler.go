package server

import (
	dom "aftp-server/pkg/domain"
	"fmt"
	"net"
)

func handleServerError(conn net.Conn) {
	message := "A server error occured"
	// TODO: calculcate byte size
	headers := []string{fmt.Sprintf("Content-Length: %d", 10)}

	createdResponse := dom.NewResponseWithContent(dom.SERVER_ERROR, headers, message)

	fmt.Printf("handleServerError with response: %s\n", createdResponse)
	doHandle(createdResponse, conn)
}
