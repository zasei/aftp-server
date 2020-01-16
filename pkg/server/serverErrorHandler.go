package server

import (
	dom "aftp-server/pkg/domain"
	"fmt"
	"net"
)

func handleServerError(conn net.Conn) {
	message := "A server error occured"
	createdResponse := dom.NewResponseWithContent(dom.SERVER_ERROR, message)

	fmt.Printf("handleServerError with response: %s\n", createdResponse)
	doHandle(createdResponse, conn)
}
