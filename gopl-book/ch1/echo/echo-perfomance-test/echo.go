package main

// intresting thing happened :  yes "arg" | head -n 100000 | xargs go run echo.go | grep elapsed was the command used to test this , but looks like xargs cannot pass all the 100000 at once , instead it is passing a set of args at a time based on the ARG_MAX of the shell !

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed for case 1\n", time.Since(start).Seconds())

	start2 := time.Now()

	var s2, sep2 string
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)
	fmt.Printf("%.2fs elapsed for case 2\n", time.Since(start2).Seconds())

	start3 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs elapsed for case 3\n", time.Since(start3).Seconds())
}
