package main

import (
	"errors"
	"fmt"
)

type Animal interface {
	name() string

	eat()
	move()
	speak()
}

type Cow struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

func (cow Cow) eat() {
	fmt.Println(cow.Food)
}

func (cow Cow) move() {
	fmt.Println(cow.Locomotion)
}

func (cow Cow) speak() {
	fmt.Println(cow.Sound)
}

func (cow Cow) name() string {
	return cow.Name
}

type Bird struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

func (bird Bird) eat() {
	fmt.Println(bird.Food)
}

func (bird Bird) move() {
	fmt.Println(bird.Locomotion)
}

func (bird Bird) speak() {
	fmt.Println(bird.Sound)
}

func (bird Bird) name() string {
	return bird.Name
}

type Snake struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

func (snake Snake) eat() {
	fmt.Println(snake.Food)
}

func (snake Snake) move() {
	fmt.Println(snake.Locomotion)
}

func (snake Snake) speak() {
	fmt.Println(snake.Sound)
}

func (snake Snake) name() string {
	return snake.Name
}

func newAnimal(name string, animalType string) (Animal, error) {

	switch animalType {
	case "cow":
		{
			return Cow{
				Name:       name,
				Food:       "grass",
				Locomotion: "walks",
				Sound:      "moo",
			}, nil
		}
	case "bird":
		{
			return Bird{
				Name:       name,
				Food:       "worms",
				Locomotion: "fly",
				Sound:      "peep",
			}, nil
		}
	case "snake":
		{
			return Snake{
				Name:       name,
				Food:       "worms",
				Locomotion: "fly",
				Sound:      "peep",
			}, nil
		}
	default:
		return nil, errors.New("no such animal type , use cow , bird or snake")
	}

}

func performAction(animal Animal, action string) {

	switch action {
	case "eat":
		animal.eat()
	case "move":
		animal.move()
	case "speak":
		animal.speak()
	default:
		fmt.Println("action not supported")
	}

}

func main() {

	fmt.Println("Enter a command")
	const initMessage = `
		usage
		: newanimal <name> <animal type>	- create new animal
		: query <name> <action>			- see animal function
		: x					- exit
    `
	fmt.Println(initMessage)

	var animals []Animal

	for {
		fmt.Print("> ")

		var action string
		var inputOne string
		var inputTwo string

		fmt.Scanf("%s %s %s", &action, &inputOne, &inputTwo)

		if action == "x" || action == "X" {
			break
		}
		if action != "query" && action != "newanimal" {
			fmt.Println("Not a valid command")
			fmt.Println(initMessage)
			continue
		}

		if len(animals) == 0 && action == "query" {
			fmt.Println("Atlest one animal must be added before querying")
			continue
		}

		if action == "newanimal" {

			if inputOne == "" || inputTwo == "" {
				fmt.Println("Animal name/type is missing")
				fmt.Println(initMessage)
				continue
			}

			animal, err := newAnimal(inputOne, inputTwo)

			if err != nil {
				fmt.Println("Could not create animal")
				continue
			}
			animals = append(animals, animal)
			fmt.Println("Created it!")
		}

		if action == "query" {

			if inputOne == "" || inputTwo == "" {
				fmt.Println("Animal name/action is missing")
				fmt.Println(initMessage)
				continue
			}

			for _, animal := range animals {

				if animal.name() == inputOne {
					performAction(animal, inputTwo)
				}

			}
		}
	}

}
