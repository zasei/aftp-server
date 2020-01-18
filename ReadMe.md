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