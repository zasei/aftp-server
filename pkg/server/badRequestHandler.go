package server

import (
	"fmt"
	"net"
)

func handleBadRequest(conn net.Conn) {
	createdResponse := Response{
		protocol:   ProtocolVersion,
		statusCode: BAD_REQUEST,
		headers:    nil,
		message:    "YOU SUCK",
	}

	fmt.Printf("handleBadRequest with response: %s\n", createdResponse)
	doHandle(createdResponse, conn)
}
