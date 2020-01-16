package server

import (
	dom "aftp-server/pkg/domain"
	"fmt"
	"net"
)

func handleBadRequest(conn net.Conn) {
	createdResponse := dom.Response{
		Protocol:   dom.ProtocolVersion,
		StatusCode: dom.BAD_REQUEST,
		Headers:    nil,
		Content:    "YOU SUCK",
	}

	fmt.Printf("handleBadRequest with response: %s\n", createdResponse)
	doHandle(createdResponse, conn)
}
