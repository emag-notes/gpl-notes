package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] == 1 {
				filenames[line] = filename
			} else if !strings.Contains(filenames[line], filename) {
				filenames[line] += " " + filename
			} else {
				// no-op
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("\t%s\n", filenames[line])
		}
	}
}
