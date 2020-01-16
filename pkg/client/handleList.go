package client

func HandleRequest(dirs []string) {
	request := Request{
		method:    LIST,
		protocol:  VERSION,
		headers:   nil,
		parameter: dirs[0],
	}

	doHandle(request)
}
