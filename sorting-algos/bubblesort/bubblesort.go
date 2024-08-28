package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var input string
	fmt.Println("Enter a sequence of numbers, separated by comma :")
	fmt.Scanln(&input)

	fmt.Println("input is : ", input)
	inputStrings := strings.Split(input, ",")

	var numbers []int

	for _, str := range inputStrings {

		num, err := strconv.Atoi(str)

		if err != nil {
			fmt.Println("Invalid input ", str)
			continue
		}
		numbers = append(numbers, num)

	}

	bubbleSort(numbers)
	fmt.Printf("the sorted array is : %v\n", numbers)
}

func bubbleSort(numbers []int) {

	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				swap(numbers, j, j+1)
			}
		}
	}

}

func swap(numbers []int, i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
