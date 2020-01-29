package server

import (
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"net"
	"os"
)

func handleDeleteRequest(request dom.Request, conn net.Conn) {
	var createdResponse dom.Response

	filePath := FileDir + request.Parameter
	md5, err := dom.HashFileMd5(filePath)
	if err != nil {
		// if error while getting hash return SERVER ERROR
		createdResponse = dom.NewResponseWithContent(dom.SERVER_ERROR, err.Error())
	} else {
		header, err := request.GetHeader(dom.ETagHeader)
		if err != nil {
			// no ETeg present in the request, return bad request
			createdResponse = dom.NewResponseWithContent(dom.BAD_REQUEST, "No ETag header present.")
		} else if md5 != header.Value {
			// hashes differ - do not delete
			createdResponse = dom.NewResponseWithContent(dom.BAD_REQUEST, "Hashes do NOT match.")
		} else {
			err := os.Remove(filePath)
			if err != nil {
				// error while removing file
				createdResponse = dom.NewResponseWithContent(dom.SERVER_ERROR, fmt.Sprintf("An error occurred while trying to remove file %s", request.Parameter))
			} else {
				// file removed successfully
				createdResponse = dom.NewResponseWithContent(dom.OK, fmt.Sprintf("File %s removed successfully", request.Parameter))
			}
		}
	}

	createdResponse.PrintResponse()
	doHandle(createdResponse, conn)
}
