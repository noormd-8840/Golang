package main

import (
	"fmt"
	"time"
)

func SendValue(c chan string){
	fmt.Println("Executing Goroutine")
	time.Sleep(1*time.Second)
	c <- "Hello"
	fmt.Println("Finished Executing Goroutine ")
}

func main(){

	values := make(chan string,2)
	defer close(values)

	go SendValue(values)
	go SendValue(values)

	valueReceived := <- values
	fmt.Println(valueReceived)

	time.Sleep(1*time.Second)
}