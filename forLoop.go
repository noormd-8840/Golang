package main
import "fmt"
func main(){

	//for loop as a while loop
	i := 1
	for i<=3{
		fmt.Println(i)
		i++
	}
	fmt.Println("Loop End \n")

	//simple for loop
	for i=0; i<10; i++{
		fmt.Println ("Execution of loop ",i)
	}
	fmt.Println("End of for loop \n")

	//for loop using  break statement
	for {
		fmt.Println("Noor Md")
		break
	}
	fmt.Println("End of for loop \n")

	//for loop using continue statement
	for i :=0; i<10; i++ {
		if i%2==0 {
			fmt.Println("Execution of if statement",i)
			continue
		}
		fmt.Println("Execution of for loop",i)
	}
}