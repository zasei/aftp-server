package client

import (
	dom "github.com/zasei/aftp-server/pkg/domain"
	"strings"
)

func HandleGetRequest(files []string) {
	request := dom.Request{
		Method:    dom.GET,
		Protocol:  dom.ProtocolVersion,
		Headers:   nil,
		Parameter: strings.Join(files, ","),
	}

	response := doHandle(request)
	response.PrintClientResponse()
}
