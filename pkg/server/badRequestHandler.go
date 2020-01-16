package server

import (
	dom "aftp-server/pkg/domain"
	"fmt"
	"net"
)

func handleBadRequest(conn net.Conn, message string) {
	createdResponse := dom.Response{
		Protocol:   dom.ProtocolVersion,
		StatusCode: dom.BAD_REQUEST,
		Headers:    nil,
		Content:    message,
	}

	fmt.Printf("handleBadRequest with response: %s\n", createdResponse)
	doHandle(createdResponse, conn)
}
