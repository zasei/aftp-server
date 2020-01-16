package client

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func HandleRequest(dirs []string) {
	request := Request{
		method:    LIST,
		protocol:  VERSION,
		headers:   nil,
		parameter: dirs[0],
	}

	doHandle(request)
}
