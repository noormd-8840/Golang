package main
import "fmt"

func plus(a int, b int) int {
	return a+b
}

func sumOfNumber(a,b,c int ) int {
	return a+b+c;
}
func main(){
	result1 := plus(1,2)
	fmt.Println("sum of two number is:",result1)

	result2 := sumOfNumber(1,2,3)
	fmt.Println("sum of three number is:",result2)
}