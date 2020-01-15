package server

import (
	"net"
)

func handleBadRequest(conn net.Conn) {
	createdResponse := Response{
		protocol:   VERSION,
		statusCode: BAD_REQUEST,
		headers:    nil,
		message:    "YOU SUCK",
	}

	doHandle(createdResponse, conn)
}
