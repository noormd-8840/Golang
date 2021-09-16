package main

import (
	"encoding/json"
	"fmt"
)

type Book struct{
	Title string 
	Author string
}

type Author struct{
	Name string
	Age int
	Developer bool
}

func main(){
	fmt.Println("Hello World!")

	book := Book{Title: "Learning cocurrency in go", Author: "Author Name"}
	author := Author{Name: "author name", Age: 30, Developer: true}
	fmt.Println("%+v\n", book)

	byteArray, err := json.MarshalIndent(book,"","  ")
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Println(byteArray)
	fmt.Println(string(byteArray))
}