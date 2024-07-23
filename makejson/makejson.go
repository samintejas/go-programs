package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	input := make(map[string]string)

	print("Enter your name : ")
	var name string
	_, errn := fmt.Scanln(&name)

	if errn != nil {
		fmt.Println("Error while reading name :", errn)
		return
	}
	input["name"] = name

	print("Enter your address : ")
	var address string
	_, erra := fmt.Scanln(&address)

	if erra != nil {
		fmt.Println("Error while reading address :", erra)
		return
	}
	input["address"] = address

	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error while marshalling JSON:", err)
		return
	}

	println("JSON : ", string(jsonData))

}
