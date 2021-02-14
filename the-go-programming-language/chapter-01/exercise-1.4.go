package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
func main() {
	counts := make(map[string]int)
	filenames := make(map[string][]string)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if !contains(filename, filenames[line]) {
				filenames[line] = append(filenames[line], filename)
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, filenames[line])
		}
	}
}

func contains(s string, array []string) bool {
	for _, e := range array {
		if e == s {
			return true
		}
	}
	return false
}
