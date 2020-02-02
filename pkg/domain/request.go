package domain

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

type Header struct {
	Name  string
	Value string
}

// Everything related to the Request - a struct with some parsing methods
type Request struct {
	Method    string
	Protocol  string
	Headers   []Header
	Parameter string
	Content   []byte
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
			requestBuilder.WriteString(fmt.Sprintf("%s: %s", s.Name, s.Value))
			requestBuilder.WriteString(NewLine)
		}
	}
	// add an empty line between request headers and body
	requestBuilder.WriteString(NewLine)

	// add the content to the request
	requestBuilder.Write(r.Content)

	// uncomment this line to view request
	//fmt.Println(requestBuilder.String())
	return requestBuilder.String()
}

func ParseRequest(requestString string) Request {

	requestLines := strings.Split(strings.TrimRight(requestString, NewLine), NewLine)

	main := strings.Fields(requestLines[0])

	var headers []Header
	var content []byte

	for i := 1; i < len(requestLines); i++ {

		if requestLines[i] == "" {
			continue
		}

		parts := strings.Split(requestLines[i], ": ")

		// if we do not have two parts, that means it was not a header line
		if len(parts) != 2 {
			content = []byte(requestLines[i])
			continue
		}

		header := Header{
			Name:  parts[0],
			Value: parts[1],
		}

		headers = append(headers, header)
	}

	parsedRequest := Request{
		Method:    main[0],
		Protocol:  main[2],
		Headers:   headers,
		Parameter: filepath.Clean(main[1]),
		Content:   content,
	}

	return parsedRequest
}

func (r Request) GetHeader(headerName string) (header Header, err error) {
	for _, header := range r.Headers {
		if strings.Contains(header.Name, headerName) {
			return header, nil
		}
	}
	return Header{}, errors.New(fmt.Sprintf("%s is not present", headerName))
}

func (r Request) PrintRequest() {
	fmt.Printf("Method: %s, Protocol: %s, Headers: %s. Parameter: %s\n", r.Method, r.Protocol, r.Headers, r.Parameter)
}
