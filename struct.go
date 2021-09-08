package main
import "fmt"

func main(){

	fmt.Println(person{"bob",20})

	fmt.Println(person{name:"Amit",age:21})

	fmt.Println(person{name:"Jhohn"})

	fmt.Println(&person{name:"Mohit",age:22})
	
	fmt.Println(newPerson("Ajay"))

	s := person{name:"Mr.Bean",age:40}
	fmt.Println(s.name)

	add := &s
	fmt.Println(add.age)

	add.age = 45
	fmt.Println(add.age)
}

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name : name}
	p.age = 24
	return &p
}