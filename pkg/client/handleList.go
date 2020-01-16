package client

import (
	dom "aftp-server/pkg/domain"
	"strings"
)

func HandleListRequest(dirs []string) {
	request := dom.Request{
		Method:    dom.LIST,
		Protocol:  dom.ProtocolVersion,
		Headers:   nil,
		Parameter: strings.Join(dirs, ","),
	}

	doHandle(request)
}
