package main
import "fmt"
func main(){

	// //array
	// //var arr [10]int = [10]int{1,2,3,4,5,6,7,8,9,0} // array
	// var arr = [...]int{1,2,3,4,5,6,7,8,9,0} //anothe way to declare and initialize array
	// fmt.Print(arr)

	//slicing 
	var arr []int = []int{1,2,3,4,5,6,7,8,9,0} //this is syntax of declaring with initialization of slicing
	fmt.Println(arr)
	fmt.Print(len(arr))

}