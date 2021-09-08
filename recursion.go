package main
import "fmt"
func main(){

	fmt.Println(fact(5))

	var fib func(n int) int
	fib = func(n int ) int {
		if n==1 || n==2 {
			return n-1
		}
		return fib(n-1) + fib(n-2) 
	}
	fmt.Println(fib(5))

}

func fact(n int) int {
	if n==0 || n==1{
		return 1
	}
	return n * fact(n-1)

}