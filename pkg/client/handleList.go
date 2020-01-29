package client

import (
	dom "github.com/zasei/aftp-server/pkg/domain"
	"strings"
)

func HandleListRequest(dirs []string) {

	since := dom.Header{
		Name:  "Since",
		Value: "1579375568",
	}

	request := dom.Request{
		Method:    dom.LIST,
		Protocol:  dom.ProtocolVersion,
		Headers:   []dom.Header{since},
		Parameter: strings.Join(dirs, ","),
	}

	response := doHandle(request)

	response.PrintClientResponse()
}
