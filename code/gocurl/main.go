package main

import (
	"fmt"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: gocurl <url>")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("In main func")
}
