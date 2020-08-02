package main

import(
"fmt"
"bufio"
"os"
"strings"
"strconv"
) 

func main() {
	strnumber:=GetInput()
	intsli:=StrtoIntSlice(strnumber)
	BubbleSort(intsli)
	
	fmt.Print(intsli)
}

func GetInput() string{
	reader:= bufio.NewReader(os.Stdin)
	var num string
	var err error

	for {
		fmt.Printf("Enter number (separate each number with a space): ")
		num, err = reader.ReadString('\n') //capture input as string
		if err == nil {
		break;
		}
		fmt.Println("Please enter correct input.")
	}
	return num
}

func StrtoIntSlice(num string) []int{

	num= strings.Trim(num, " \r\n")
	numarr:=strings.Split(num, " ") //split all of the numbers
	intsli:=make([]int,0,10)

	for i:=0; i< len(numarr); i++ {
		number, err:= strconv.Atoi(numarr[i]) //turn each string into int

		if err!= nil{
			fmt.Println("Only integers are accepted") 
			fmt.Print(err)
		break;
		}
		
		intsli=append(intsli,number) //append the int number onto the slice
	}
	return intsli
}

func BubbleSort(intsli []int){
	var contSortflag bool //loop again the sorting loop if true
	for{
		contSortflag=false
		for i:=0; i< (len(intsli)-1); i++ { //sorting loop
			if intsli[i]>intsli[i+1]{
				Swap(intsli,i) //swap places if intsli[i]>intsli[i+1]
				contSortflag=true
			}
		}
		if contSortflag == false{
			break;
		}
	}
}

func Swap(intsli []int,i int){
	tempVar:=intsli[i]	
	intsli[i]=intsli[i+1]
	intsli[i+1]=tempVar
}
