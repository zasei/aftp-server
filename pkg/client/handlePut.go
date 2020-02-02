package client

import (
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func HandlePutRequest(files []string, etags []string) {
	parameters := strings.Join(files, ",")
	for i, filePath := range files {
		// open the file and get the info
		file, err := os.Open(filePath)
		check(err, filePath)

		fileInfo, err := file.Stat()
		check(err, filePath)

		headers := make([]dom.Header, 1)
		// create the last modified header from timestamp
		lastModified := dom.Header{
			Name:  dom.LastModifiedHeader,
			Value: strconv.FormatInt(fileInfo.ModTime().Unix(), 10),
		}
		headers = append(headers, lastModified)

		// add the etag header if present
		if len(etags) != 0 {
			etagHeader := dom.Header{
				Name:  dom.ETagHeader,
				Value: etags[i],
			}
			headers = append(headers, etagHeader)
		}

		// read the contents from it
		dat, err := ioutil.ReadFile(filePath)
		check(err, filePath)

		request := dom.Request{
			Method:    dom.PUT,
			Protocol:  dom.ProtocolVersion,
			Headers:   headers,
			Parameter: parameters,
			Content:   dat,
		}

		response := doHandle(request)
		response.PrintClientResponse()
	}
}

func check(err error, path string) {
	if err != nil {
		fmt.Printf("Unable to read/open file from %s\n", path)
		os.Exit(1)
	}
}
