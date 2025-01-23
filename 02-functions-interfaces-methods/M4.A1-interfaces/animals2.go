package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func newCow() Cow {
	return Cow{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
}

func (c Cow) Eat()   { fmt.Printf("%s\n", c.food) }
func (c Cow) Move()  { fmt.Printf("%s\n", c.locomotion) }
func (c Cow) Speak() { fmt.Printf("%s\n", c.noise) }

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func newBird() Bird {
	return Bird{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
}

func (b Bird) Eat()   { fmt.Printf("%s\n", b.food) }
func (b Bird) Move()  { fmt.Printf("%s\n", b.locomotion) }
func (b Bird) Speak() { fmt.Printf("%s\n", b.noise) }

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func newSnake() Snake {
	return Snake{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
}

func (s Snake) Eat()   { fmt.Printf("%s\n", s.food) }
func (s Snake) Move()  { fmt.Printf("%s\n", s.locomotion) }
func (s Snake) Speak() { fmt.Printf("%s\n", s.noise) }

func main() {
	const (
		prompt = ">"

		create = "newanimal"
		query  = "query"

		cow   = "cow"
		bird  = "bird"
		snake = "snake"

		eat   = "eat"
		move  = "move"
		speak = "speak"
	)

	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%s ", prompt)

		if ok := scanner.Scan(); !ok {
			fmt.Fprintf(os.Stderr, "\n\nerror: %v\n", scanner.Err())
			continue
		}

		input := strings.Fields(scanner.Text())
		cmd := input[0]

		switch cmd {
		case create:
			name := input[1]
			typeOf := input[2]

			switch typeOf {
			case cow:
				animals[name] = newCow()
			case bird:
				animals[name] = newBird()
			case snake:
				animals[name] = newSnake()
			default:
				fmt.Fprintf(os.Stderr, "invalid animal: %s\n", typeOf)
			}

		case query:
			name := input[1]
			behavior := input[2]

			if animal, ok := animals[name]; ok {
				switch behavior {
				case eat:
					animal.Eat()
				case move:
					animal.Move()
				case speak:
					animal.Speak()
				default:
					fmt.Fprintf(os.Stderr, "invalid behavior: %s\n", behavior)
				}
			} else {
				fmt.Fprintf(os.Stderr, "invalid name: %s\n", name)
			}
		default:
			fmt.Fprintf(os.Stderr, "invalid command: %s\n", cmd)
		}
	}
}
