package client

import (
	dom "aftp-server/pkg/domain"
	"strings"
)

func HandleGetRequest(files []string) {
	request := dom.Request{
		Method:    dom.GET,
		Protocol:  dom.ProtocolVersion,
		Headers:   nil,
		Parameter: strings.Join(files, ","),
	}

	doHandle(request)
}
