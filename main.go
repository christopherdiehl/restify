package main

import (
	"fmt"
	"os"

	"github.com/christopherdiehl/restify/api"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide specify location of file")
		os.Exit(1)
	}
	api.CreateServer(os.Args[1])
	// start up server here
}
