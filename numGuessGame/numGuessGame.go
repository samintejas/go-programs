package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	banner := `
	

   _____                     __  __       _   _                 _               
  / ____|                   |  \/  |     | \ | |               | |              
 | |  __ _   _  ___  ___ ___| \  / |_   _|  \| |_   _ _ __ ___ | |__   ___ _ __ 
 | | |_ | | | |/ _ \/ __/ __| |\/| | | | | .   | | | | '_   _ \| '_ \ / _ \ '__|
 | |__| | |_| |  __/\__ \__ \ |  | | |_| | |\  | |_| | | | | | | |_) |  __/ |   
  \_____|\__,_|\___||___/___/_|  |_|\__, |_| \_|\__,_|_| |_| |_|_.__/ \___|_|   
                                     __/ |                                      
                                    |___/                                       


	`
	fmt.Println(banner)

	// generate a random number between 1 and 100
	var randomNumber int
	var userInput int
	var userInputRaw string
	tries := 0

	randomNumber = rand.Intn(100)

	fmt.Println(":: Hi! ive guessed a number between 0 and 100")
	fmt.Println(":: Guess the number")

	for {

		tries++

		fmt.Print(" > ")
		_, err := fmt.Scanln(&userInputRaw)
		userInput, err = strconv.Atoi(userInputRaw)

		if err != nil || userInput < 0 || userInput > 100 {
			fmt.Println(":: Hmm.. Thats not a valid, try again !")
			continue
		}

		if userInput == randomNumber {
			fmt.Printf(":: U've won !! u tried %d times", tries)
			break
		} else if userInput < randomNumber {
			fmt.Println(":: nah.. the number is larger")
		} else {
			fmt.Println(":: nah.. the number is smaller")
		}

	}

}
