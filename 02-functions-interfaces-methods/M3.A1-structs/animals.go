package main

import (
	"fmt"
	"os"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat()   { fmt.Println(a.food) }
func (a Animal) Move()  { fmt.Println(a.locomotion) }
func (a Animal) Speak() { fmt.Println(a.noise) }

func main() {
	const prompt = ">"

	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	for {
		fmt.Printf("%s ", prompt)

		var animalStr, behaviorStr string
		_, err := fmt.Scan(&animalStr, &behaviorStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n\nerror: %v\n", err)
			os.Exit(1)
		}

		animal, ok := animals[animalStr]
		if !ok {
			fmt.Fprintf(os.Stderr, "unknown animal\n\n")
			continue
		}

		switch behaviorStr {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Fprintf(os.Stderr, "unknown behavior\n\n")
		}
	}
}
