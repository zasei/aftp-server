package server

import (
	"fmt"
	"net"
	"strings"
)

const VERSION = "AFTP/1.0"
const FILE_DIR = "files"

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

func HandleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	//Read the incoming connection into the buffer
	_, err := conn.Read(buf)

	requestString := strings.Fields(string(buf))

	if err != nil {
		fmt.Println("Error reading", err.Error())
		handleServerError(conn)
	}

	if len(requestString) < 3 {
		handleBadRequest(conn)
	}
	parsedRequest := parseRequest(requestString)
	switch parsedRequest.method {
	case GET:
		handleGetRequest(parsedRequest, conn)
	case LIST:
		handleListRequest(parsedRequest, conn)
	default:
		handleBadRequest(conn)
	}
}

func doHandle(response Response, conn net.Conn) {
	_, _ = conn.Write([]byte(createResponse(response)))
	err := conn.Close()
	if err != nil {
		fmt.Printf("Error sending response %s", err)
	}
}

type Request struct {
	method     string
	protocol   string
	headers    []string
	parameters []string
}

func parseRequest(requestString []string) Request {
	parseRequest := Request{
		method:   requestString[0],
		protocol: requestString[2],
		headers:  nil,
		// TODO: Parameter parsing logic
		parameters: []string{requestString[1]},
	}

	return parseRequest
}

type Response struct {
	protocol   string
	statusCode string
	headers    []string
	message    string
}

func createResponse(response Response) string {
	var responseString string
	if len(response.headers) == 0 {
		responseString = response.protocol + " " + response.statusCode + "\n" + strings.Join(response.headers, "\n") + "\n\n" + response.message
	} else if len(response.message) != 0 {
		responseString = response.protocol + " " + response.statusCode + "\n" + response.message
	} else {
		responseString = response.protocol + " " + response.statusCode
	}
	return responseString
}
