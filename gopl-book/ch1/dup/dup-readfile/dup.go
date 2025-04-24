package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, fileName := range os.Args[1:] {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup-readfile %v \n", err)
			continue
		}

		// used splitSeq here because it is more efficient
		for line := range strings.SplitSeq(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d : %s \n", n, line)
		}
	}
}
