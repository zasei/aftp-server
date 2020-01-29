package client

import (
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"os"
	"strings"
)

func HandleDeleteRequest(files []string) {
	// loop over all parameters and handle each separately
	for _, f := range files {
		// check the file - try to open, if not there exit
		openFile, err := os.Open(strings.TrimLeft(f, "/"))
		if err != nil {
			fmt.Println("No file found to create hash from - please make sure the file exists locally before removing")
			os.Exit(1)
		}

		// get hash of the file
		md5, _ := dom.HashFileMd5(openFile.Name())
		fmt.Println("md5: " + md5)

		etagHeader := dom.Header{
			Name:  dom.ETagHeader,
			Value: md5,
		}

		// create the request
		request := dom.Request{
			Method:    dom.DELETE,
			Protocol:  dom.ProtocolVersion,
			Headers:   []dom.Header{etagHeader},
			Parameter: strings.Join(files, ","),
		}

		request.PrintRequest()

		response := doHandle(request)
		response.PrintClientResponse()
	}
}
