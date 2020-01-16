package server

import (
	"fmt"
	"net"
)

func handleServerError(conn net.Conn) {
	createdResponse := Response{
		protocol:   ProtocolVersion,
		statusCode: SERVER_ERROR,
		headers:    nil,
		message:    "YOU SUCK",
	}

	fmt.Printf("handleServerError with response: %s\n", createdResponse)
	doHandle(createdResponse, conn)
}
