package main

import (
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

const maxNameLength = 20

func main() {

	fmt.Print("Enter the name of the text file (place file in current directory) : ")
	var filePath string
	fmt.Scanln(&filePath)

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}

	lines := strings.Split(string(content), "\n")

	var names []Name

	for _, line := range lines {
		// Trim spaces and skip empty lines
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		fname, lname := parts[0], parts[1]
		if len(fname) > maxNameLength {
			fname = fname[:maxNameLength]
		}
		if len(lname) > maxNameLength {
			lname = lname[:maxNameLength]
		}

		name := Name{
			fname: fname,
			lname: lname,
		}
		names = append(names, name)
	}

	fmt.Println("\nNames from file:")
	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.fname, name.lname)
	}
}
