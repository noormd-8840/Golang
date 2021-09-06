package main
import "fmt"
func main(){
	fmt.Println("enter value")
	var myString string
 
	//this is for single input 
	fmt.Scanln(&myString)
	fmt.Println(myString)
}