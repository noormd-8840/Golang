package main
import "fmt"

func main(){

	var obj Vehicle = Bike{Name: "Honda",Colour: "Black",Price: 100000}

	obj.ShowDetails()

	name := obj.ShowName()
	fmt.Println(name)

}

type Vehicle interface {
	ShowDetails()
	ShowName() string
}

type Bike struct {
	Name,Colour string
	Price float64
}

func (bike Bike) ShowDetails() {
	fmt.Println("Bike Name: ",bike.Name)
	fmt.Println("Bike Colour: ",bike.Colour)
	fmt.Println("Bike Price: ",bike.Price)
}

func (bike Bike) ShowName() string{
	return bike.Name
}


