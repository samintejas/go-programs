package main

import "fmt"

// to write a unicode character in neovim , go to insert mode -> use ctrl+v -> type u followed by a unicode (03b1 for alpha symbol)
func main() {
	fmt.Println("Hello, α")
}
