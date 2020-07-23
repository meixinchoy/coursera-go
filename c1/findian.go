package main

import(
"fmt"
"strings"
) 

func main() {
	var str string
	fmt.Printf("Enter string: ")
	fmt.Scan(&str)
	str= strings.ToLower(str)

	if strings.HasPrefix(str, "i") && str[len(str)-1] == 'n' && strings.Contains(str,"a"){
	 	fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}