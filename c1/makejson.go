package main

import(
"fmt"
"bufio"
"encoding/json"
"os"
"strings"
) 

func main() {
	var err error
	var name string
	var address string
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Scan()
	for {
		fmt.Printf("Enter name: ")
		name, err = reader.ReadString('\n')
		if err == nil {
		break;
		}
		fmt.Println("Please enter correct input.\n")
	}
	
	for {
		fmt.Printf("Enter address: ")
		address, err = reader.ReadString('\n')
		if err == nil {
		break;
		}
		fmt.Println("Please enter correct input.\n")
	}
	
	map1 := make(map[string] string)
	map1["name"]=strings.Trim(name, " \r\n")
	map1["address"]=strings.Trim(address, " \r\n")

	
	for {
		barr, err := json.Marshal(map1)
		if err == nil {
		fmt.Printf("%s", barr)
		break
		}
	}	
}