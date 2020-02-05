package server

import (
	"fmt"
	dom "github.com/zasei/aftp-server/pkg/domain"
	"net"
	"os"
)

func handlePutRequest(request dom.Request, conn net.Conn) {
	// TODO: Implement

	var createdResponse dom.Response
	filePath := FileDir + "/" + request.Parameter

	etagHeader, err := request.GetHeader(dom.ETagHeader)
	if err == nil {
		// ETag etagHeader is present, handle flow which puts over EXISTING file
		fmt.Printf("ETag etagHeader present: %s\n", etagHeader.Value)

		// check if file exists - if it does not but ETag is given - return 404
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			createdResponse = dom.NewResponseNotFound()
			doHandle(createdResponse, conn)
			// return to ensure the other code flows do not run
			return
		}

		// open file and check hashes
		file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
		if err != nil {
			createdResponse = internalServerError(filePath)
		} else if file != nil {
			md5, _ := dom.HashFileMd5(filePath)

			if md5 != etagHeader.Value {
				// md5 hashes do NOT match - return GONE status
				createdResponse = dom.NewResponseWithContent(dom.GONE, "md5 hashes do not match")
			} else {
				// md5 hashes DO match - OVERWRITE file
				_, err := file.Write(request.Content)
				md5, _ := dom.HashFileMd5(filePath)
				createdResponse = dom.NewResponseWithContent(dom.OK, md5)
				if err != nil {
					createdResponse.PrintResponse()
					createdResponse = internalServerError(filePath)
				}
			}
		}

		createdResponse.PrintResponse()
		doHandle(createdResponse, conn)
	}

	// ETag etagHeader is NOT present, handle uploading of NEW file

	// check if file already exists - if it does we need the ETag header
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File does not exist - we do NOT need the ETag header and we can create the new file
		// create file
		file, err := os.Create(filePath)
		if err != nil {
			// handle error of being unable to create file
			createdResponse = internalServerError(request.Parameter)
		} else if file != nil {
			// file is present, write to it
			_, err := file.Write(request.Content)
			md5, err := dom.HashFileMd5(filePath)
			if err != nil {
				createdResponse = internalServerError(request.Parameter)
			}
			createdResponse = dom.NewResponseWithContent(dom.OK, md5)
		}
	} else {
		// file does exist -  we need the ETag header
		createdResponse = dom.NewResponseWithContent(dom.LOCKED, "md5 hash is missing")
	}

	createdResponse.PrintResponse()
	doHandle(createdResponse, conn)
}

func internalServerError(filePath string) dom.Response {
	errorMessage := fmt.Sprintf("Could not open, create or read file %s\n", filePath)
	fmt.Println(errorMessage)
	return dom.NewResponseWithContent(dom.SERVER_ERROR, errorMessage)

}
