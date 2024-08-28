package main

import "fmt"

func main() {

	var a string = "type specified"
	var b = "type infered"
	var c, d = 1, 2
	e := "shorthand for initialization"
	var f = true
	var g int

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c,d", c, d)
	fmt.Println("e", e)
	fmt.Println("f", f)
	// all non initialized variables in go will have 0 value
	fmt.Println("g", g)

}
