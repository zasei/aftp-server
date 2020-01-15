package client

const VERSION = "AFTP/1.0"

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
	BAD_REQUEST  = "400 Bad request"
	NOT_FOUND    = "404 Not found"
	GONE         = "418 Gone"
	LOCKED       = "423 Locked"
	SERVER_ERROR = "500 Server Error"
)

const (
	HOST = "127.0.0.1"
	PORT = "1337"
	TYPE = "tcp"
)

func createRequest(request Request) string {
	var requestString string
	if len(request.headers) == 0 {
		if len(request.parameter) == 0 {
			requestString = request.method + " " + request.protocol
		} else {
			requestString = request.method + " " + request.parameter + " " + request.protocol
		}
	}
	return requestString
}
