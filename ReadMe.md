# AFTP Server and Client in GO Lang

## Project setup
Ensure that you have this project in your $GOPATH/src directory
Ensure that you change the FileDir constant in serverHandler.go to the path of your choosing.

### Work with Client
Install extra packages:
```bash
go get github.com/ogier/pflag
```

### Run/build
To run and/or build, follow these steps:
```bash
# Enter the dir of the client
cd cmd/aftp

# This installs the binary locally
go install

# After which, you can run it like this:
go run aftp-client.go
# Or running it from the command line, as it is installed as a binary on your path (YAY Go)
# Running this will show the command options
aftp

# Or build the binary like this:
go build

# After which, the binary is created and placed in the dir and can be executed:
# Running this will show the command options
aftp
```



# AFTP file protocol
The protocol is state-less. It therefore does not matter which requests the client sent earlier in the session. New lines consist of \ r \ n. All text except the content is UTF-8 encoded.

## Status codes
The status code always consists of a three-digit code followed by the textual description of that code.
The possible status codes are:
200 OK - The request has been successfully executed. The response contains the requested answer
400 Bad request - The request does not comply with the protocol. The server does not understand what the intention is.
404 Not found - The request cannot be executed. The requested data cannot be found.
418 Gone - The file in the specified version does not exist.
423 Locked - The file is described as by another client.
500 Server Error - The request was good, but something is not going right in the server, so the request cannot be answered.

## Headers
Headers are extra information about requests and responses. The keys are case insensitive. The value is case sensitive depending on the key. A header always consists of a key followed by a colon and a space followed by the value of the header.

## UNIX timestamp
The UNIX timestamp is a standard format in UNIX to indicate date time. The timestamp is a number consisting of the number of seconds since 1 January 1970 at midnight.

## Requests
The first line of a request consists of the command, a space and a relative path. This path always starts with a forward slash (/). The line then continues with a space, the word AFTP followed by a forward slash (/) followed by the version. Currently only version 1.0.
The following lines follow the headers, one on each line. The headers end when an empty line is sent.
Any content then follows after that blank line. The amount of content is equal to the value of the Content-Length bytes header.

## Responses
A response consists of the word AFTP followed by a forward slash (/) followed by the version. Currently only version 1.0. The line is continued with the status code. The following lines follow the response headers followed by an empty line. Optionally, content will follow after the amount of content is defined in the Content-Length header. The header indicates how many bytes the content is.

## Request file list
The complete file list can be requested with the LIST command. This happens when a new application is connected. The path can be supplemented to limit the response to the specified path. Usually the application will only ask for the changes since its last request. This is done by giving the Since header followed by the UNIX timestamp from when changes must be returned.
In both cases, the content of the responses to these requests are a list of file names followed by a space and a UNIX timestamp when this file was last modified, followed by a space and the md5 hash of the content of the file. One file per line.

## Request file
A file can be retrieved with the GET command and the path of the file.
The content of the response will be the content of the file. If the file is not found, a response with status code 404 is returned. The response will then not contain any content.

## Create / edit file
A file can be created or modified by the PUT command. The file name is in the path. The Last-Modified header indicates when the file was last modified as a UNIX timestamp. The content of the request is filled with the content of the file. If the file needs to be modified, the ETag header must be included. The value of this header contains the previous md5 hash known to the client of the content of the file.
The content of the response is successfully filled with the md5 hash of the content of the received file. The response code indicates whether it was successful.

## Delete file
A file can be deleted by the DELETE command. The file name is in the path. The ETag header must be included. The value of this header contains the previous md5 hash known to the client of the content of the file.
The response code indicates whether it was successful.

## Example
\>LIST / AFTP/1.0  
\>  
\<AFTP/1.0 200 OK  
\<Content-Length: xxx  
\<  
\<Tekst.txt 1575800743 25d7f36115f74a2fa7c6d185ab008588  
\<Afbeelding.jpg 1575800657 99f4f1eba9e7ce89a7d5b2e8631dcfc5  
\>GET /Tekst.txt AFTP/1.0  
\>  
\<AFTP/1.0 200 OK  
\<Content-Length: xxx  
\<  
\<RGl0IGlzIGVlbiB0ZWtzdCBiZXN0YW5kCk1ldAoKZWVuIGFhbnRhbCAKCm5ldwpsaW5lcwoK  
\>PUT /Tekst.txt AFTP/1.0  
\>Content-Length: xxx  
\>ETag: 25d7f36115f74a2fa7c6d185ab008588  
\>  
\>RGl0IGlzIGVlbiBhbmRlciB0ZWtzdCBiZXN0YW5kCk1ldAoKZWVuIGFhbnRhbCAKCm5ldwpsaW5lcwoK  
\<AFTP/1.0 200 OK  
\<Content-Length: xxx  
\<  
\<b02c4735bc5bd845f94b0e8b515c3303  
\>LIST / AFTP/1.0  
\>Since: 1575800744  
\>  
\<AFTP/1.0 200 OK  
\<Content-Length: xxx  
\<  
\<Tekst.txt 1575800744 b02c4735bc5bd845f94b0e8b515c3303
