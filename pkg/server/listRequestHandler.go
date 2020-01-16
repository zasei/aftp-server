package server

import (
	dom "aftp-server/pkg/domain"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

func handleListRequest(request dom.Request, conn net.Conn) {
	content := listDirectory(request.Parameter)

	var createdResponse dom.Response
	if len(content) == 0 || strings.Contains(content, "no files found") {
		createdResponse = dom.NewResponseNotFound()
	} else {
		createdResponse = dom.NewResponseWithContent(dom.OK, content)
	}

	createdResponse.PrintResponse()
	doHandle(createdResponse, conn)
}

func listDirectory(path string) string {
	//files, err := ioutil.ReadDir("./" + FileDir + path)
	files, err := ioutil.ReadDir(FileDir + path)

	if err != nil {
		return fmt.Sprintf("no files found for %s", path)
	}

	var results strings.Builder

	for _, f := range files {
		results.WriteString(fmt.Sprintf("%s %d %s\n", f.Name(), f.ModTime().Unix(), "MD5 HERE"))
	}

	return results.String()
}
