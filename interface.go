package main

import (
	"fmt"
	"math"
)

func main(){
	r := rectangle{width: 5,hieght: 5}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

type geometric interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width,hieght float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width + r.hieght
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.hieght
} 

func (c circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func (c circle) perimeter() float64 {
	return 2*math.Pi*c.radius
}

func measure(g geometric){

	fmt.Println(g)

	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}