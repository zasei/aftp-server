package domain

import (
	"fmt"
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
	if len(r.Headers) == 0 {
		if len(r.Parameter) == 0 {
			//fmt.Println("0 headers, 0 parameters")
			return fmt.Sprintf("%s %s\r\n", r.Method, r.Protocol)
		} else {
			//fmt.Println("0 headers, >1 parameters")
			return fmt.Sprintf("%s %s %s", r.Method, r.Parameter, r.Protocol)
		}
	} else {
		if len(r.Parameter) == 0 {
			//fmt.Println(">1 headers, 0 parameters")
			return fmt.Sprintf("%s %s\r\n%s", r.Method, r.Protocol, r.Headers)
		} else {
			//fmt.Println(">1 headers, >1 parameters")
			return fmt.Sprintf("%s %s %s\r\n%s", r.Method, r.Parameter, r.Protocol, r.Headers)
		}
	}
}

func ParseRequest(requestString []string) Request {
	parseRequest := Request{
		Method:    requestString[0],
		Protocol:  requestString[2],
		Headers:   nil,
		Parameter: requestString[1],
	}

	return parseRequest
}

func (r Request) printRequest() {
	fmt.Printf("Method: %s, Protocol: %s, Headers: %s. Parameter: %s", r.Method, r.Protocol, r.Headers, r.Parameter)
}
