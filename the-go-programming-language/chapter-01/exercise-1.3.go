package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Exercise 1.3: Experiment to measure the difference in running time between
// our potentially inefficient versions and the one that uses strings.Join.
func main() {
	start1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	end1 := time.Now()
	fmt.Println(end1.Sub(start1))

	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end2 := time.Now()
	fmt.Println(end2.Sub(start2))

}
