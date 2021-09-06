package main

import "fmt"

func main(){
	empSal := map[string]int{
		"Neha" : 100,
		"Raj" : 200,
		"Anaya" : 300,
	}

		//ankit, ok := empSal["Ankit"]
		//ankit, ok := empSal["Neha"] // to check availability
		_,flag := empSal["Neha"]//instead of _ we use code later
		//ankit,flag := empSal["Neha"]
		//fmt.Println(ankit,flag)
		fmt.Println(flag)
}

