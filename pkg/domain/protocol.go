package domain

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

const (
	ProtocolVersion = "AFTP/1.0"
	Separator       = " "
	NewLine         = "\r\n"
	ETagHeader      = "ETAG:"
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
	OK           = "200 OK"
	BAD_REQUEST  = "400 Bad Request"
	NOT_FOUND    = "404 Not Found"
	GONE         = "418 Gone"
	LOCKED       = "423 Locked"
	SERVER_ERROR = "500 Server Error"
)

const (
	HOST = "127.0.0.1"
	PORT = "1337"
	TYPE = "tcp"
)

func CalculateContentLength(content string) int {
	if len(content) == 0 {
		return 0
	}
	// convert the string to a byte array, a byte in GoLang is always 1 byte aka 8 bit
	// https://golang.org/ref/spec#Size_and_alignment_guarantees
	return 8 * len([]byte(content))
}

func HashFileMd5(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}
