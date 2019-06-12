package main

import (
	"fmt"
	"os"
)
func main() {

	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hellow world")
	}


	os.Exit(-1)
}