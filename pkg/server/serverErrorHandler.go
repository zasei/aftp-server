package server

import (
	"net"
)

func handleServerError(conn net.Conn) {
	createdResponse := Response{
		protocol:   VERSION,
		statusCode: SERVER_ERROR,
		headers:    nil,
		message:    "YOU SUCK",
	}

	doHandle(createdResponse, conn)
}
