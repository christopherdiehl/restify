package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide specify location of file")
		os.Exit(1)
	}
	CreateServer(os.Args[1])
	// start up server here
}
