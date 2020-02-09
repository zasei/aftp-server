package client

import (
	"bytes"
	"encoding/gob"
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"io"
	"log"
	"net"
	"os"
)

// Handle the actual network connection and parsing
func doHandle(request dom.Request) dom.Response {
	// Set up connection to remote
	conn, connErr := net.Dial(dom.TYPE, dom.HOST+":"+dom.PORT)
	if connErr != nil {
		fmt.Printf("Error while connecting %s\n", connErr)
		os.Exit(1)
	}

	defer conn.Close()

	enc := gob.NewEncoder(conn)

	err := enc.Encode(request)

	if err != nil {
		log.Fatal("encode error:", err)
	}

	var buf bytes.Buffer
	io.Copy(&buf, conn)
	receivedResponse := dom.ParseResponse(buf)

	return receivedResponse
}
