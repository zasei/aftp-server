package server

import (
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
)

func handleListRequest(request dom.Request, conn net.Conn) {

	sinceHeader, _ := request.GetHeader(dom.SinceHeader)

	content := listDirectory(request.Parameter, sinceHeader)

	var createdResponse dom.Response

	if len(content) == 0 || strings.Contains(content, "no files found") {
		createdResponse = dom.NewResponseNotFound()
	} else {
		createdResponse = dom.NewResponseWithContent(dom.OK, content)
	}

	createdResponse.PrintResponse()
	doHandle(createdResponse, conn)
}

func listDirectory(path string, header dom.Header) string {

	files, err := ioutil.ReadDir(FileDir + path)

	if err != nil {
		return fmt.Sprintf("no files found for %s", path)
	}

	var results strings.Builder

	for _, f := range files {

		md5 := ""

		if !f.IsDir() {
			md5, _ = dom.HashFileMd5(FileDir + path + f.Name())
		}

		if header.Value != "" {

			sinceTime, _ := strconv.ParseInt(header.Value, 10, 64)

			timestamp := f.ModTime().Unix()

			fmt.Println(timestamp)

			// 150600000 >= 150699
			if f.ModTime().Unix() >= sinceTime {
				results.WriteString(fmt.Sprintf("%s %d %s\n", f.Name(), uint64(f.ModTime().Unix()), md5))
			}

		} else {
			results.WriteString(fmt.Sprintf("%s %d %s\n", f.Name(), uint64(f.ModTime().Unix()), md5))
		}

	}

	return results.String()
}
