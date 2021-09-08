package main
import "fmt"
func main(){
	sum(1,2)
	sum(1,2,3,4,5)

	nums := []int{5,4,3,2,1}
	sum(nums...)
}

func sum(nums ...int){
	fmt.Println(nums, " ")
	total := 0
	for _,number := range nums {
		total += number
	}
	fmt.Println(total)

}