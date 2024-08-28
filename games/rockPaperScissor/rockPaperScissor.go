package main

import (
	"fmt"
	"math/rand"
)

func main() {

    fmt.Println("-- rock paper scissor --")

    sysGuess := map[int]string {
        1: "rock",
        2: "paper",
        3: "scissor",
    }

    for {

        randomNum := rand.Intn(3) + 1

        fmt.Printf("The random guess is %d \n", randomNum)
        var userInput int
        fmt.Printf("Stone .. Papper .. Scissor !!\n")
        fmt.Printf("[1] rock \n [2] paper \n [3] scissor\n")
        fmt.Scanf("%d",&userInput)

        if sysGuess[userInput] == sysGuess[randomNum] {
            fmt.Println("same !!") 
        }

        if sysGuess[userInput] == "rock" && sysGuess[randomNum] == "paper" {
            fmt.Println("You lost")
        }

        if sysGuess[userInput] == "rock" && sysGuess[randomNum] == "scissor" {
            fmt.Println("You won")
        }

        if sysGuess[userInput] == "paper" && sysGuess[randomNum] == "scissor" {
            fmt.Println("You lost")
        }

        if sysGuess[userInput] == "paper" && sysGuess[randomNum] == "stone" {
            fmt.Println("You won")
        }

        if sysGuess[userInput] == "scissor" && sysGuess[randomNum] == "rock" {
            fmt.Println("You lost")
        }

        if sysGuess[userInput] == "scissor" && sysGuess[randomNum] == "paper" {
            fmt.Println("You won")
        }




    }

}
