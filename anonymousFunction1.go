package main
import "fmt"

func main(){

	// //Anonymous function without return type
	// f := func ()  {
	// 	fmt.Println("Hello, All")
	// }

	// f()

	// //with return type
	// f := func () string {
	// 	return "Hi"
	// }

	// fmt.Println(f())

	//Anonymous fuction with immediate call 
	func ()  {
		fmt.Println("Hello")
	}()
}