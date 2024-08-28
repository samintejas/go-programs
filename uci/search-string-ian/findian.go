package main

import("fmt")
import("strings")


func main() {
	
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	input = strings.TrimSpace(input)
	lowerCased := strings.ToLower(input)

	startsWithI := strings.HasPrefix(lowerCased,"i")
	endsWithN := strings.HasSuffix(lowerCased,"n") 
	containsA := strings.Contains(lowerCased,"a")

	if  startsWithI && endsWithN && containsA {	
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
