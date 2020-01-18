package client

import (
	dom "github.com/zasei/aftp-server/pkg/domain"
	"strings"
)

func HandleListRequest(dirs []string) {
	request := dom.Request{
		Method:    dom.LIST,
		Protocol:  dom.ProtocolVersion,
		Headers:   nil,
		Parameter: strings.Join(dirs, ","),
	}

	response := doHandle(request)

	response.PrintClientResponse()
}
