package client

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const VERSION = "AFTP/1.0"

// request options
const (
	GET    = "GET"
	LIST   = "LIST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// response options
const (
	OK           = "200 OK"
	BAD_REQUEST  = "400 Bad request"
	NOT_FOUND    = "404 Not found"
	GONE         = "418 Gone"
	LOCKED       = "423 Locked"
	SERVER_ERROR = "500 Server Error"
)

const (
	HOST = "127.0.0.1"
	PORT = "1337"
	TYPE = "tcp"
)

// Handle the actual network connection and parsing
func doHandle(request Request) {
	fmt.Println("doHandle")
	// Set up connection to remote
	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Printf("Error while connecting %s\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	// Create the request as a String literal
	requestString := request.requestToString()

	// Write to remote connection
	fmt.Fprintf(conn, requestString)

	// create buffer and copy bytes to it
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	receivedResponse := parseResponse(buf)

	receivedResponse.printResponse()
}

// Everything related to the Request - a struct with some parsing methods
type Request struct {
	method    string
	protocol  string
	headers   []string
	parameter string
}

func (request Request) requestToString() string {
	var requestString string
	if len(request.headers) == 0 {
		if len(request.parameter) == 0 {
			requestString = request.method + " " + request.protocol
		} else {
			requestString = request.method + " " + request.parameter + " " + request.protocol
		}
	}
	return requestString
}

// Everything related to the Response - a struct with some parsing methods
type Response struct {
	protocol   string
	statusCode string
	headers    []string
	message    string
}

func parseResponse(buf bytes.Buffer) Response {
	// convert byte buffer to string
	responseString := buf.String()
	splitResponse := strings.Split(responseString, "\r\n")

	splitFirstLine := strings.Split(splitResponse[0], " ")
	receivedResponse := Response{
		protocol: splitFirstLine[0],
	}

	if strings.Contains(responseString, OK) {
		receivedResponse.statusCode = OK
	} else if strings.Contains(responseString, BAD_REQUEST) {
		receivedResponse.statusCode = BAD_REQUEST
	} else if strings.Contains(responseString, NOT_FOUND) {
		receivedResponse.statusCode = NOT_FOUND
	} else if strings.Contains(responseString, GONE) {
		receivedResponse.statusCode = GONE
	} else if strings.Contains(responseString, LOCKED) {
		receivedResponse.statusCode = LOCKED
	} else if strings.Contains(responseString, SERVER_ERROR) {
		receivedResponse.statusCode = SERVER_ERROR
	} else {
		fmt.Println("Unable to get status code from response")
		os.Exit(1)
	}

	// TODO: figure out how to split rest properly
	if len(splitResponse) == 2 {
		receivedResponse.message = splitResponse[1]
	}

	return receivedResponse
}

func (r Response) printResponse() {
	fmt.Printf("Response: { protocol: %s, statusCode: %s, headers: %s, message: %s } \n", r.protocol, r.statusCode, r.headers, r.message)
}
