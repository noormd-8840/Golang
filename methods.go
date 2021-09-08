package main
import "fmt"

type rectangle struct {
	width,hieght int
}

func (r *rectangle) area() int {
	return r.width*r.hieght
}

func (r rectangle) perim() int {
	return 2*r.width + 2*r.hieght
}

func main(){
	r := rectangle{width: 10,hieght: 20}

	fmt.Println("area: ",r.area())
	fmt.Println("perimeter: ",r.perim())

	rp := &r
	fmt.Println("area: ",rp.area())
	fmt.Println("perimeter: ",rp.perim())
}