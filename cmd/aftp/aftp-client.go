package main

import (
	handler "aftp-server/pkg/client"
	"fmt"
	flag "github.com/ogier/pflag"
	"os"
	"strings"
)

var (
	ls string
	cp string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(ls) != 0 {
		dirs := strings.Split(ls, ",")
		fmt.Printf("Listing files for: %s\n", dirs)
		handler.HandleRequest(dirs)
	}

}

func init() {
	flag.StringVarP(&ls, "list", "l", "", "List files")
	flag.StringVarP(&cp, "copy", "c", "", "Copy file")
}
