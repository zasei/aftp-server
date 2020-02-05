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
	response.Headers = []string{fmt.Sprintf("%s: %d", ContentLengthHeader, CalculateContentLength(content))}
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

	splitResponse := strings.Split(responseString, NewLine)

	splitFirstLine := strings.Split(splitResponse[0], Separator)
	receivedResponse := Response{
		Protocol: splitFirstLine[0],
	}

	if receivedResponse.Protocol != ProtocolVersion {
		fmt.Println("Protocols do not match, exiting")
		os.Exit(1)
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
	if len(splitResponse) == 3 {
		receivedResponse.Headers = make([]string, 1)
		receivedResponse.Headers[0] = splitResponse[1]
	}
	if len(splitResponse) == 4 {
		receivedResponse.Headers = make([]string, 1)
		receivedResponse.Headers[0] = splitResponse[1]
		// splitResponse[2] is an empty line
		receivedResponse.Content = splitResponse[3]
	}

	//fmt.Printf("Showing split response lines, total length: %d\n", len(splitResponse))
	//for _, s := range splitResponse {
	//	fmt.Println(s)
	//}

	return receivedResponse
}

func (r Response) CreateResponse() string {
	// create stringBuilder
	var responseBuilder strings.Builder
	// add protocol version and start a new line
	responseBuilder.WriteString(r.Protocol)
	responseBuilder.WriteString(Separator)
	// add status code
	responseBuilder.WriteString(r.StatusCode)
	responseBuilder.WriteString(NewLine)
	// if headers are present, add them
	if len(r.Headers) != 0 {
		for _, s := range r.Headers {
			responseBuilder.WriteString(s)
			responseBuilder.WriteString(NewLine)
		}
		// add an empty line between request headers and body
		responseBuilder.WriteString(NewLine)
	}
	// add content if present
	if len(r.Content) != 0 {
		fmt.Printf("Adding content: %s\n", r.Content)
		responseBuilder.WriteString(r.Content)
	}
	// return stringbuilder as string
	fmt.Println(responseBuilder.String())
	return responseBuilder.String()
}

func (r Response) PrintClientResponse() {
	switch r.StatusCode {
	case OK:
		if len(r.Content) == 0 {
			fmt.Printf("%s, %s", r.StatusCode, r.Headers)
		} else {
			fmt.Println(removeNewLine(r.Content))
		}
	default:
	case BAD_REQUEST, NOT_FOUND, SERVER_ERROR, LOCKED:
		if len(r.Content) == 0 {
			fmt.Printf("Received status code %s", removeNewLine(r.StatusCode))
		} else {
			fmt.Printf("Received status code %s with message: %s", removeNewLine(r.StatusCode), r.Content)
		}
	}
}

func (r Response) PrintResponse() {
	fmt.Printf("Response: { Protocol: %s, StatusCode: %s, Headers: %s, Content: %s } \n", r.Protocol, r.StatusCode, r.Headers, r.Content)
}

func removeNewLine(msg string) string {
	return strings.TrimRight(msg, NewLine)
}
