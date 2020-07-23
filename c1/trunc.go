package main

import "fmt"

func trunc() {
	var num float32
	fmt.Printf("Enter floating point number: ")
	fmt.Scan(&num)
	fmt.Printf("Integer number: ")
	fmt.Println(int(num))
}