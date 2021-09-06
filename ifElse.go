package main
import "fmt"
func main() {

	fmt.Println("Enter number")
    var i int 
	fmt.Scanln(&i)
	if i%2 == 0 {
		fmt.Println(i, "is even number")
	} else {
		fmt.Println(i, "is odd number")
	}

	//greatest number among three number
	a, b, c := 5, 4, 3

	if a > b && a > c {
		fmt.Println("a is greatest")
	} else if b > a && b > c {
		fmt.Println("b is greatest")
	} else {
		fmt.Println("c is greatest")
	}

}