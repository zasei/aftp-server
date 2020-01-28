package client

import (
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"os"
	"strings"
)

func HandleGetRequest(files []string) {
	request := dom.Request{
		Method:    dom.GET,
		Protocol:  dom.ProtocolVersion,
		Headers:   nil,
		Parameter: strings.Join(files, ","),
	}

	response := doHandle(request)

	for _, f := range files {
		filePath := strings.TrimLeft(f, "/")
		fmt.Println("Creating local file " + filePath)
		createdFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error while creating local file " + filePath)
			fmt.Println(err)
		}
		createdFile.WriteString(response.Content)
	}
	response.PrintClientResponse()
}
