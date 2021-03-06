package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	handler "github.com/zasei/aftp-server/pkg/client"
	"os"
	"strings"
)

var (
	list   string
	get    string
	put    string
	etag   string
	remove string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(list) != 0 {
		dirs := strings.Split(list, ",")
		fmt.Printf("Listing files for: %s\n", dirs)
		handler.HandleListRequest(dirs)
	} else if len(get) != 0 {
		files := strings.Split(get, ",")
		fmt.Printf("")
		handler.HandleGetRequest(files)
	} else if len(put) != 0 {
		files := strings.Split(put, ",")
		fmt.Printf("")
		if len(etag) != 0 {
			// ETag given
			handler.HandlePutRequest(files, strings.Split(etag, ","))
		} else {
			// NO ETag given
			handler.HandlePutRequest(files, make([]string, 0))
		}
	} else if len(remove) != 0 {
		files := strings.Split(remove, ",")
		fmt.Printf("")
		handler.HandleDeleteRequest(files)
	}
}

func init() {
	flag.StringVarP(&list, "list", "l", "", "List files")
	flag.StringVarP(&get, "get", "g", "", "Copy file")
	flag.StringVarP(&put, "put", "p", "", "Put a file")
	flag.StringVarP(&etag, "etag", "e", "", "ETag to provide for PUT")
	flag.StringVarP(&remove, "remove", "r", "", "Remove a file")
}
