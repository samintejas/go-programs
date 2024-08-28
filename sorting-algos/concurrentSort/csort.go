package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {


    var arr []int

    for {
        var inputnum string
        fmt.Println("enter the number you want to sort ( one at a time pls ) , press X to sort")
        fmt.Scanf("%s",&inputnum)


        if(inputnum == "X" || inputnum == "x") {
            break
        }

        num,err := strconv.Atoi(inputnum)

        if err != nil {
            fmt.Println("not a valid number")
            continue
        }

        arr = append(arr,num)
    }

    length := len(arr)
    quater := length/4

    var wg sync.WaitGroup
    wg.Add(4)

    sortedParts := make([][]int,4)

    for i := 0; i <4;i++ {
        start := i * quater
        end := start + quater
        if i ==3 {
            end = length
        }
        sortedParts[i] = arr[start:end]
        go func(i int) {
            defer wg.Done()
            quicksort(sortedParts[i])
        }(i)
    }

    wg.Wait()

    result := merge(sortedParts[0],sortedParts[1],sortedParts[2],sortedParts[3])
    fmt.Println("The sorted numbers are : ",result)
}

func merge(arrs ...[]int) []int {
	result := arrs[0]
	for _, arr := range arrs[1:] {
		result = mergeTwo(result, arr)
	}
	return result
}

func mergeTwo(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}
func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	left := []int{}
	right := []int{}

	for i, v := range arr {
		if i == pivotIndex {
			continue
		}
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	quicksort(left)
	quicksort(right)

	copy(arr, append(append(left, pivot), right...))
}

