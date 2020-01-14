package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
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
	OK = "200 OK"
	BAD_REQUEST = "400 Bad request"
	NOT_FOUND = "404 Not found"
	GONE = "418 Gone"
	LOCKED = "423 Locked"
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

func handleRequest(conn net.Conn) {

	buf := make([]byte, 1024)

	//Read the incoming connection into the buffer
	_, err := conn.Read(buf)

	requestString := strings.Fields(string(buf))

	fmt.Println(requestString[0])

	var response string

	switch requestString[0] {
	case GET:
		getFile(requestString[1])
	case LIST:
		response = listDirectory(requestString[1])
	default:
		response = BAD_REQUEST + "\n\n"
	}

	if err != nil {
		fmt.Println("Error reading", err.Error())
		returnResponse(SERVER_ERROR, conn)
	}
	_, _ = conn.Write([]byte(response))

	_ = conn.Close()

}


func returnResponse(response_code string, conn net.Conn) {
	_, _ =conn.Write([]byte(response_code))
}