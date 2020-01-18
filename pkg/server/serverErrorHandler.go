package server

import (
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"net"
)

func handleServerError(conn net.Conn) {
	message := "A server error occured"
	createdResponse := dom.NewResponseWithContent(dom.SERVER_ERROR, message)

	fmt.Printf("handleServerError with response: %s\n", createdResponse)
	doHandle(createdResponse, conn)
}
