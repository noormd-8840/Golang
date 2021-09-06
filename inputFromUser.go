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
	fmt.Println("Enter a Number")
	reader := bufio.NewReader(os.Stdin)
	input,err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error while reading input")
	}

	trimmedInput := strings.TrimSpace(input)
	n,err := strconv.ParseInt(trimmedInput,10,64)
	if err != nil {
		log.Fatal("Input was not a number")
	}
	fmt.Println(n)
}