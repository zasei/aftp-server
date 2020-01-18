package server

import (
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func handleGetRequest(request dom.Request, conn net.Conn) {
	var createdResponse dom.Response

	filesData := getFiles(request.Parameter)
	if len(filesData) == 0 {
		createdResponse = dom.NewResponseNotFound()
	} else {
		content := strings.Join(filesData, "\r\n")
		createdResponse = dom.NewResponseWithContent(dom.OK, content)
	}

	createdResponse.PrintResponse()
	doHandle(createdResponse, conn)
}

func getFiles(parameter string) []string {
	// get all files from parameter
	filePaths := strings.Split(parameter, ",")

	nrOfFiles := len(filePaths)

	filesData := make([][]byte, nrOfFiles)
	filesContent := make([]string, nrOfFiles)
	// loop over the files
	for i, s := range filePaths {
		path := FileDir + s
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			continue
		}
		fmt.Println(string(dat))
		filesData[i] = dat
		filesContent[i] = string(dat)

		f, err := os.Open(path)
		check(err)

		b1 := make([]byte, 32)
		n1, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

		err = f.Close()
		check(err)
	}
	return filesContent
}
