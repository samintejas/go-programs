package main

import (
	"fmt"
)

type Animal struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

func (animal Animal) eat() {
	fmt.Printf("%s eat : %s \n", animal.Name, animal.Food)
}

func (animal Animal) move() {
	fmt.Printf("%s move : %s \n", animal.Name, animal.Locomotion)
}

func (animal Animal) speak() {
	fmt.Printf("%s speak : %s \n", animal.Name, animal.Sound)
}

func main() {

	cow := Animal{
		Name:       "cow",
		Food:       "grass",
		Locomotion: "walk",
		Sound:      "moo",
	}

	bird := Animal{
		Name:       "bird",
		Food:       "worms",
		Locomotion: "fly",
		Sound:      "peep",
	}
	snake := Animal{
		Name:       "snake",
		Food:       "mice",
		Locomotion: "slither",
		Sound:      "hss",
	}

	const initMessage = `
    Enter a command
    usage : <name> <action>  - to see the action assosiated with the animal
          : x                - exit
    `

	fmt.Println(initMessage)

	for {

		fmt.Print("> ")
		var animalName string
		var animalAction string

		fmt.Scanf("%s %s", &animalName, &animalAction)

		parts := []string{animalName, animalAction}

		if animalName == "\n" || animalName == "" {
			continue
		}
		if parts[0] == "x" || parts[0] == "X" {
			break
		}

		switch parts[0] {
		case "cow":
			{
				doAction(cow, parts[1])
			}
		case "bird":
			{

				doAction(bird, parts[1])
			}
		case "snake":
			{

				doAction(snake, parts[1])
			}
		default:
			fmt.Printf("%s is not a predefined animal \nvalid options are : \n\tcow\n\tbird\n\tsnake\n", parts[0])
		}

	}
}

func doAction(animal Animal, action string) {

	switch action {
	case "eat":
		animal.eat()
	case "move":
		animal.move()
	case "speak":
		animal.speak()
	case "":
		fmt.Println("Please enter an action")
	default:
		fmt.Printf("%s is not a pre defined action, \nvalid options are : \n\teat\n\tspeak\n\tmove\n", action)
	}

}
