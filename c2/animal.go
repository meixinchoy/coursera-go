package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct {
	name, food, locomotion, sound string
}

func (a *Animal) SetAttributes(n,f,l,s string){
	a.name=n
	a.food=f
	a.locomotion=l
	a.sound=s
}

func (a Animal) Name() string{
	return a.name
}

func (a Animal) Move(){
	fmt.Println(a.locomotion)
}

func (a Animal) Eat(){
	fmt.Println(a.food)
}

func (a Animal) Speak(){
	fmt.Println(a.sound)
}

func main(){
	var animals[3] Animal
	animals[0].SetAttributes("cow","grass","walk","moo")
	animals[1].SetAttributes("bird","worms","fly","peep")
	animals[2].SetAttributes("snake","mice","slither","hsss")
	reader := bufio.NewReader(os.Stdin)
	var request string
	var err error

	for {
		var checking bool = false
		fmt.Print("\n>")
		request, err = reader.ReadString('\n')
		request= strings.Trim(request,"\n\r")
		requestArr:= strings.Split(request, " ")
		

		if err == nil {
			for i:=0; i<3; i++ {
				if requestArr[0]==animals[i].Name(){
					if requestArr[1]=="eat"{
						animals[i].Eat()
					}else if requestArr[1]=="speak"{
						animals[i].Speak()
					}else if requestArr[1]=="move"{
						animals[i].Move()
					}else{
						checking=false
						break;
					}
					checking=true
					break
				}
			}
		}
		if checking== false{
			fmt.Print("Please enter valid input.\n")
		}
	}
}