package main

import(
"fmt"
"os"
"strings"
"bufio"
) 

type Name struct {
	fname string
	lname string
}

func main() {
var fileName string
fmt.Print("Enter name of text file: ")
fmt.Scan(&fileName)

f,err:= os.Open(fileName)
if err!= nil{
	fmt.Println("error opening file")
	return
}

scanner:= bufio.NewScanner(f)
sli:= make([]Name,0)
for scanner.Scan(){
	line:= scanner.Text() //get names from file line by line
	names := strings.Split(line, " ") //separate fname and lname n put into names[2] 
	n:=Name{fname:names[0],lname:names[1]} //create a struct for each name
	sli=append(sli,n) //append the struct onto the slice
}

fmt.Println(sli)

for i:=0; i<len(sli);i++{
	fmt.Println(sli[i].fname+" "+sli[i].lname)
} 
}