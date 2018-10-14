package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	// name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	// name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	// name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var command, name, subcommand string
	var animalDb map[string]Animal
	animalDb = make(map[string]Animal)
	for {

		fmt.Print("> ")
		fmt.Scan(&command)
		fmt.Scan(&name)
		fmt.Scan(&subcommand)

		switch command {
		case "newanimal":
			if _, ok := animalDb[name]; ok {
				fmt.Printf("Name '%s' already exists\n", name)
			} else {
				switch subcommand {
				case "cow":
					animalDb[name] = Cow{}
					fmt.Println("created it!")
				case "bird":
					animalDb[name] = Bird{}
					fmt.Println("created it!")
				case "snake":
					animalDb[name] = Snake{}
					fmt.Println("created it!")
				default:
					fmt.Printf("invalid animal - '%s'\n", subcommand)
				}
			}

		case "query":
			animal, ok := animalDb[name]
			if !ok {
				fmt.Printf("name '%s' doesnt exist in animal db!\n", name)
			} else {
				switch subcommand {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Printf("invalid action - '%s'\n", subcommand)
				}
			}

		default:
			fmt.Printf("Invalid command - %s\n", command)
		}

		fmt.Println()
	}
}
