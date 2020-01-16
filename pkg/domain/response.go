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
	Message    string
}

func ParseResponse(buf bytes.Buffer) Response {
	// convert byte buffer to string
	responseString := buf.String()
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
		fmt.Println("Unable to get status code from response")
		os.Exit(1)
	}

	// TODO: figure out how to split rest properly
	if len(splitResponse) == 2 {
		receivedResponse.Message = splitResponse[1]
	}

	return receivedResponse
}

func CreateResponse(response Response) string {
	if len(response.Headers) != 0 {
		if len(response.Message) == 0 {
			return response.Protocol + " " + response.StatusCode + "\r\n" + strings.Join(response.Headers, "\r\n") + "\r\n" + response.Message
		} else {
			return response.Protocol + " " + response.StatusCode + "\r\n" + strings.Join(response.Headers, "\r\n") + "\r\n"
		}
	}
	if len(response.Message) != 0 {
		return response.Protocol + " " + response.StatusCode + "\r\n" + response.Message
	} else {
		return response.Protocol + " " + response.StatusCode
	}
}

func (r Response) PrintResponse() {
	fmt.Printf("Response: { Protocol: %s, StatusCode: %s, Headers: %s, Message: %s } \n", r.Protocol, r.StatusCode, r.Headers, r.Message)
}
