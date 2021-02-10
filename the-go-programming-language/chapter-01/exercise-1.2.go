package main

import (
	"fmt"
	"os"
	"strconv"
)

// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
func main() {
	for index, args := range os.Args[1:] {
		fmt.Println(strconv.Itoa(index+1) + " " + args)
	}
}
