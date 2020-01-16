package server

import (
	dom "aftp-server/pkg/domain"
	"fmt"
	"net"
)

func handleServerError(conn net.Conn) {
	createdResponse := dom.Response{
		Protocol:   dom.ProtocolVersion,
		StatusCode: dom.SERVER_ERROR,
		Headers:    nil,
		Message:    "YOU SUCK",
	}

	fmt.Printf("handleServerError with response: %s\n", createdResponse)
	doHandle(createdResponse, conn)
}
