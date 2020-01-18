package domain

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Everything related to the Request - a struct with some parsing methods
type Request struct {
	Method    string
	Protocol  string
	Headers   []string
	Parameter string
}

func (r Request) RequestToString() string {
	// uncomment this line to view the full request struct, before parsing it to a string
	//r.printRequest()
	// create stringBuilder
	var requestBuilder strings.Builder
	// first, add the Method and a space
	requestBuilder.WriteString(r.Method)
	requestBuilder.WriteString(Separator)
	// if parameter is present, add it and a space
	if len(r.Parameter) != 0 {
		requestBuilder.WriteString(r.Parameter)
		requestBuilder.WriteString(Separator)
	}
	// add protocol version and start a new line
	requestBuilder.WriteString(r.Protocol)
	requestBuilder.WriteString(NewLine)
	// if headers are present, add them
	if len(r.Headers) != 0 {
		for _, s := range r.Headers {
			requestBuilder.WriteString(s)
			requestBuilder.WriteString(NewLine)
		}
	}
	// add an empty line between request headers and body
	requestBuilder.WriteString(NewLine)
	// uncomment this line to view request
	//fmt.Println(requestBuilder.String())
	return requestBuilder.String()
}

func ParseRequest(requestString []string) Request {
	parseRequest := Request{
		Method:    requestString[0],
		Protocol:  requestString[2],
		Headers:   nil,
		Parameter: filepath.Clean(requestString[1]),
	}

	return parseRequest
}

func (r Request) PrintRequest() {
	fmt.Printf("Method: %s, Protocol: %s, Headers: %s. Parameter: %s", r.Method, r.Protocol, r.Headers, r.Parameter)
}
