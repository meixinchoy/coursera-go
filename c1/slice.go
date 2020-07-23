package main

import(
"fmt"
"strconv"
"strings"
"sort"
) 

func main() {
	sli:= make([]int, 0,3)
	var numstr string

	for{
		fmt.Printf("Enter int: ")
		fmt.Scan(&numstr)
		num, err:= strconv.Atoi(numstr)
		if strings.ToUpper(numstr) == "X"{
			break
		} else if err==nil{
			sli= append(sli,num)
			sort.Slice(sli, func(i,j int) bool{ return sli[i] < sli[j]})
			fmt.Println(sli)
		}else{
			fmt.Println("Invalid integer, please enter again.")
		}
	}
	fmt.Println("Program ended")
}