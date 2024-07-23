package main

import (
	"fmt"
	"strconv"
)

func main() {

	numbers := make([]int, 0, 10)

	fmt.Println("Press X to exit")

	for {

		var input string

		fmt.Print("Enter a number [or press x or X to exit] : ")
		_, err := fmt.Scanln(&input)

		number, err := strconv.Atoi(input)

		if err != nil {
			if input == "X" || input == "x" {
				break
			} else {
				fmt.Println("> Please enter a valid number")
				continue
			}
		}

		numbers = append(numbers, number)

		sort(numbers)
		fmt.Println("sorted array : ", numbers)

	}

	fmt.Println("Exiting program !")

}

func sort(numberSlice []int) {

	n := len(numberSlice)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if numberSlice[j] > numberSlice[j+1] {
				numberSlice[j], numberSlice[j+1] = numberSlice[j+1], numberSlice[j]
			}
		}
	}

}
