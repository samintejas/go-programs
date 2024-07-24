package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("-- Displacement calculator --")

	var genDisp func(time float64) float64

	acceleration := askUser("Acceleration")
	initVel := askUser("initial Velocity")
	initDisp := askUser("initial Displacement")

	genDisp = GenDisplaceFn(acceleration, initVel, initDisp)
	var isFirstTime bool = true
	for {

		if isFirstTime {
			printDisplacement(genDisp)
			isFirstTime = false
		}

		fmt.Print("Do you want to calculate displacement for another time [y/n] : ")
		var userInp string
		fmt.Scanln(&userInp)

		if userInp == "y" || userInp == "Y" {
			printDisplacement(genDisp)
		} else if userInp == "n" || userInp == "N" {
			break
		} else {
			fmt.Println("Acceptable input y : yes , n : no")
		}

	}

}

func printDisplacement(genDisp func(time float64) float64) {
	time := askUser("Time")
	fmt.Printf("the displacement is : %f \n", genDisp(time))
}

func askUser(keyWord string) float64 {
	for {
		fmt.Printf("Enter %s : ", keyWord)
		var userInp string
		fmt.Scanln(&userInp)

		floatValue, err := strconv.ParseFloat(userInp, 64)

		if err != nil {
			fmt.Println("not a valid acceleration value (decimal numbers are accepted)")
			continue
		}
		return floatValue
	}
}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(time float64) float64 {

	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + s0
	}
}
