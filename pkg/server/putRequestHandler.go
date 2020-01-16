package server

import (
	dom "aftp-server/pkg/domain"
	"net"
)

func handlePutRequest(request dom.Request, conn net.Conn) {
	// TODO: Implement
	createdResponse := dom.NewResponseWithContent(dom.BAD_REQUEST, "method not yet implemented")

	createdResponse.PrintResponse()
	doHandle(createdResponse, conn)
}
