package domain

const (
	ProtocolVersion = "AFTP/1.0"
	Separator       = " "
	NewLine         = "\r\n"
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
