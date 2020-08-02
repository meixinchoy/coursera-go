package main

import(
"fmt"
"bufio"
"os"
"strings"
"strconv"
) 

func main() {
	reader:= bufio.NewReader(os.Stdin)
	var num string
	var err error
	numsli:=make([]int,0,10)

	for {
		fmt.Printf("Enter number (separate each number with a space): ")
		num, err = reader.ReadString('\n') //capture input as string
		if err == nil {
		break;
		}
		fmt.Println("Please enter correct input.")
	}
	num= strings.Trim(num, " \r\n")
	numarr:=strings.Split(num, " ") //split all of the numbers

	for i:=0; i< len(numarr); i++ {
		number, err:= strconv.Atoi(numarr[i]) //turn each string into int

		if err!= nil{
			fmt.Println("Only integers are accepted") 
			fmt.Print(err)
		break;
		}
		
		numsli=append(numsli,number) //append the int number onto the slice
	}

	var contSortflag bool //loop again the sorting loop if true
	for{
		contSortflag=false
		for i:=0; i< (len(numarr)-1); i++ { //sorting loop
			if numsli[i]>numsli[i+1]{
			tempVar:=numsli[i]	//swap places if numsli[i]>numsli[i+1]
			numsli[i]=numsli[i+1]
			numsli[i+1]=tempVar
			contSortflag=true
			}
		}
		if contSortflag == false{
			break;
		}
	}

	fmt.Print(numsli)

}