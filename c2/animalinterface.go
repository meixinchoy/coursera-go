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
	name string
}

func (a Cow) Eat()   { fmt.Println("grass") }
func (a Cow) Move()  { fmt.Println("walk") }
func (a Cow) Speak() { fmt.Println("moo") }

type Bird struct {
	name string
}

func (a Bird) Eat()   { fmt.Println("worms") }
func (a Bird) Move()  { fmt.Println("fly") }
func (a Bird) Speak() { fmt.Println("peep") }

type Snake struct {
	name string
}

func (a Snake) Eat()   { fmt.Println("mice") }
func (a Snake) Move()  { fmt.Println("slither") }
func (a Snake) Speak() { fmt.Println("hsss") }

func main() {
	m := make(map[string]Animal) //create map to store animals
	reader := bufio.NewReader(os.Stdin)
	var input string
	var err error

	for {
		var success bool = false
		fmt.Print("\n>")
		input, err = reader.ReadString('\n')
		input = strings.Trim(input, "\n\r")
		inputArr := strings.Split(input, " ") //split input into several strs

		if err == nil && len(inputArr) != 3 {
			name := inputArr[1]
			if inputArr[0] == "newanimal" { //if first key word == newanimal
				var a Animal
				success = true       //initialize success as true
				switch inputArr[2] { //switch animal type
				case "cow":
					a = Cow{name}
				case "snake":
					a = Snake{name}
				case "bird":
					a = Bird{name} //create objects
				default:
					success = false //success=false if type doest match
				}
				m[name] = a //put new obj into the map
				if success {
					fmt.Println("Created it!") //print msg if obj created successfully
				}
			} else if inputArr[0] == "query" {
				for key, val := range m { //for all items in the map
					if name == key { //correct animal name
						if inputArr[2] == "eat" {
							val.Eat() //eat ins
						} else if inputArr[2] == "speak" {
							val.Speak() //speak ins
						} else if inputArr[2] == "move" {
							val.Move() //move ins
						} else {
							success = false // incorrect insructions, success=false
							break
						}
						success = true // instructions executed, success=true
					}
				} //if incorrect name, success is false by default
			}
			//success=false by default (if keyword newanimal or query wasnt entered)
		}
		if success == false { // display err msg
			fmt.Print("Please enter valid input.\n")
		}
	} //program continue to loop forever
}

//peer review ans:
/*
package main

import (
	"fmt"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (a animal) Eat() {
	fmt.Println(a.food)
	return
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
	return
}

func (a animal) Speak() {
	fmt.Println(a.noise)
	return
}

func main() {
	animalMap := make(map[string]animal)
	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worms", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hsss"}
	var genralAni animalInterface
	for {
		var command, requestAni, requestType string
		fmt.Print(">")
		fmt.Scan(&command, &requestAni, &requestType)
		if command == "query" {
			genralAni = animalMap[requestAni]
			switch requestType {
			case "eat":
				genralAni.Eat()
			case "move":
				genralAni.Move()
			case "speak":
				genralAni.Speak()
			}
		}
		if command == "newanimal" {
			animalMap[requestAni] = animalMap[requestType]
			fmt.Println("Created it!")
		}
	}
}
*/
