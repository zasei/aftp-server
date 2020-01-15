package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const FILE_DIR = "files"
const VERSION = "AFTP/1.0"

// server settings
const (
	HOST = "127.0.0.1"
	PORT = "1337"
	TYPE = "tcp"
)

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

func main() {

	l, err := net.Listen(TYPE, HOST+":"+PORT)

	if err != nil {
		fmt.Printf("Error listening on %s:%s", HOST, PORT)
		os.Exit(1)
	}

	defer l.Close()

	fmt.Printf("Listening on %s:%s", HOST, PORT)

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting new connection", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine
		go handleRequest(conn)
	}

}

func getFile(path string) {
	fmt.Printf("Get file %s", path)

}

func deleteFile(path string) {

}

func putFile(path string) {

}

func listDirectory(path string) string {
	files, err := ioutil.ReadDir("./" + FILE_DIR + path)

	if err != nil {
		log.Fatal(err)
	}

	var results strings.Builder

	for _, f := range files {
		results.WriteString(f.Name() + "\n")
	}

	return results.String()
}

type Response struct {
	protocol   string
	statusCode string
	headers    []string
	message    string
}

func createResponse(response Response) string {
	responseString := response.protocol + " " + response.statusCode + "\n" + strings.Join(response.headers, "\n") + "\n\n" + response.message
	return responseString
}

type Request struct {
	method     string
	protocol   string
	headers    []string
	parameters []string
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	//Read the incoming connection into the buffer
	_, err := conn.Read(buf)

	requestString := strings.Fields(string(buf))
	println(requestString)

	if err != nil {
		fmt.Println("Error reading", err.Error())
		handleServerError(conn)
	}

	if len(requestString) < 4 {
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

func handleGetRequest(request Request, conn net.Conn) {

}

func handleListRequest(request Request, conn net.Conn) {
	response := Response{
		protocol:   VERSION,
		statusCode: OK,
		headers:    nil,
		message:    listDirectory(request.parameters[0]),
	}

	fmt.Print(response)
	doHandle(response, conn)
}

func parseRequest(requestString []string) Request {
	parseRequest := Request{
		method:   requestString[0],
		protocol: requestString[2],
		headers:  nil,
		// TODO: Parameter parsing logic
		parameters: []string{requestString[1]},
	}

	fmt.Println(parseRequest)
	return parseRequest
}

func handleBadRequest(conn net.Conn) {
	createdResponse := Response{
		protocol:   VERSION,
		statusCode: BAD_REQUEST,
		headers:    nil,
		message:    "YOU SUCK",
	}

	doHandle(createdResponse, conn)
}

func handleServerError(conn net.Conn) {
	createdResponse := Response{
		protocol:   VERSION,
		statusCode: SERVER_ERROR,
		headers:    nil,
		message:    "YOU SUCK",
	}

	doHandle(createdResponse, conn)
}

func doHandle(response Response, conn net.Conn) {
	_, _ = conn.Write([]byte(createResponse(response)))
	conn.Close()
}

func returnResponse(responseCode string) string {
	var content = ""
	var headers []string

	// TODO: FIX HEADERS
	if content != "" {
		headers[0] = "Content-Length: " + strconv.Itoa(len(content))
	} else {
		headers[0] = "Content-Length: 0"
	}
	return VERSION + " " + responseCode + "\n" + strings.Join(headers, "\n")
}
