package client

import (
	dom "aftp-server/pkg/domain"
	"strings"
)

func HandlePutRequest(files []string) {
	// TODO: Implement
	request := dom.Request{
		Method:    dom.PUT,
		Protocol:  dom.ProtocolVersion,
		Headers:   nil,
		Parameter: strings.Join(files, ","),
	}

	response := doHandle(request)
	response.PrintClientResponse()
}
