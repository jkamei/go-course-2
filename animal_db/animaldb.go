package main

import ("fmt"
)

type Animal struct {
		food string
		locomotion string
		noise string
	}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hiss"}

	for {
		var animalInput, actionInput string

		fmt.Printf("> ")
		fmt.Scan(&animalInput)
		fmt.Scan(&actionInput)

		switch animalInput {
			case "cow":
				switch actionInput {
					case "eat": cow.Eat()
					case "move": cow.Move()
					case "speak": cow.Speak()
				}
			case "bird":
				switch actionInput {
					case "eat": bird.Eat()
					case "move": bird.Move()
					case "speak": bird.Speak()
				}
			case "snake":
				switch actionInput {
					case "eat": snake.Eat()
					case "move": snake.Move()
					case "speak": snake.Speak()
				}
		}
	}
}