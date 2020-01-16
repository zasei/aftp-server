package client

import (
	dom "aftp-server/pkg/domain"
	"strings"
)

func HandleDeleteRequest(files []string) {
	// TODO: Implement
	request := dom.Request{
		Method:    dom.DELETE,
		Protocol:  dom.ProtocolVersion,
		Headers:   nil,
		Parameter: strings.Join(files, ","),
	}

	response := doHandle(request)
	response.PrintClientResponse()
}
