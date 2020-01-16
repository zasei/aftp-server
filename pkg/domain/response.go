package domain

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Everything related to the Response - a struct with some parsing methods
type Response struct {
	Protocol   string
	StatusCode string
	Headers    []string
	Content    string
}

func NewResponseWithContent(statusCode string, content string) Response {
	response := Response{}
	response.Protocol = ProtocolVersion
	response.StatusCode = statusCode
	response.Headers = []string{fmt.Sprintf("Content-Length: %d", CalculateContentLength(content))}
	response.Content = content
	return response
}

func NewResponseNotFound() Response {
	response := Response{}
	response.StatusCode = NOT_FOUND
	response.Content = ""
	response.Protocol = ProtocolVersion
	return response
}

func ParseResponse(buf bytes.Buffer) Response {
	// convert byte buffer to string
	responseString := buf.String()

	// uncomment this line to show the response before parsing
	//fmt.Printf("Response string received: %s\n", responseString)

	splitResponse := strings.Split(responseString, "\r\n")

	splitFirstLine := strings.Split(splitResponse[0], " ")
	receivedResponse := Response{
		Protocol: splitFirstLine[0],
	}

	if strings.Contains(responseString, OK) {
		receivedResponse.StatusCode = OK
	} else if strings.Contains(responseString, BAD_REQUEST) {
		receivedResponse.StatusCode = BAD_REQUEST
	} else if strings.Contains(responseString, NOT_FOUND) {
		receivedResponse.StatusCode = NOT_FOUND
	} else if strings.Contains(responseString, GONE) {
		receivedResponse.StatusCode = GONE
	} else if strings.Contains(responseString, LOCKED) {
		receivedResponse.StatusCode = LOCKED
	} else if strings.Contains(responseString, SERVER_ERROR) {
		receivedResponse.StatusCode = SERVER_ERROR
	} else {
		fmt.Println("Unable to get status code from response, got: " + responseString)
		os.Exit(1)
	}

	// TODO: figure out how to split rest properly
	if len(splitResponse) == 2 {
		receivedResponse.Content = splitResponse[1]
	}
	if len(splitResponse) == 3 {
		receivedResponse.Headers = make([]string, 1)
		receivedResponse.Headers[0] = splitResponse[1]
		receivedResponse.Content = splitResponse[2]
	}

	//fmt.Printf("Showing split response lines, total length: %d\n", len(splitResponse))
	//for _, s := range splitResponse {
	//	fmt.Println(s)
	//}

	return receivedResponse
}

func (r Response) CreateResponse() string {
	if len(r.Headers) != 0 {
		if len(r.Content) == 0 {
			return r.Protocol + " " + r.StatusCode + "\r\n" + strings.Join(r.Headers, "\r\n") + "\r\n"
		} else {
			return r.Protocol + " " + r.StatusCode + "\r\n" + strings.Join(r.Headers, "\r\n") + "\r\n" + r.Content
		}
	}
	if len(r.Content) != 0 {
		return r.Protocol + " " + r.StatusCode + "\r\n" + r.Content
	} else {
		return r.Protocol + " " + r.StatusCode + "\r\n"
	}
}

func (r Response) PrintClientResponse() {
	switch r.StatusCode {
	case OK:
		fmt.Println(r.Content)
	default:
	case BAD_REQUEST, SERVER_ERROR:
		fmt.Printf("%s\n%s", r.StatusCode, r.Content)
	}
}

func (r Response) PrintResponse() {
	fmt.Printf("Response: { Protocol: %s, StatusCode: %s, Headers: %s, Content: %s } \n", r.Protocol, r.StatusCode, r.Headers, r.Content)
}
