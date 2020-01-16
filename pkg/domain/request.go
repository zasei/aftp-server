package domain

import "strings"

// Everything related to the Request - a struct with some parsing methods
type Request struct {
	Method    string
	Protocol  string
	Headers   []string
	Parameter string
}

func (request Request) RequestToString() string {
	var requestString string
	if len(request.Headers) == 0 {
		if len(request.Parameter) == 0 {
			requestString = request.Method + " " + request.Protocol
		} else {
			requestString = request.Method + " " + request.Parameter + " " + request.Protocol
		}
	}
	return requestString
}

func ParseRequest(requestString []string) Request {
	parameters := []string{requestString[1], requestString[2]}
	parseRequest := Request{
		Method:   requestString[0],
		Protocol: requestString[2],
		Headers:  nil,
		// TODO: Parameter parsing logic
		Parameter: strings.Join(parameters, ","),
	}

	return parseRequest
}
