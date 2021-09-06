package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func main(){

	fmt.Println("enter value")
	reader := bufio.NewReader(os.Stdin)
	input,err := reader.ReadString('\n')
	if err != nil{
		log.Fatal("Error while reading input")
	}

	trimmedInput := strings.TrimSpace(input)
	value,err:= strconv.ParseInt(trimmedInput,10,32)

	if err != nil{
		log.Fatal("Input was not a number")
	}

	switch value {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thrusday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("this is default and its execute when there is no cases are matched")
	}
}