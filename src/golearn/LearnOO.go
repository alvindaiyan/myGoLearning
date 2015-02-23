package main

import "fmt"
import "math"

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func area(r Rectangle) float64 {
	return r.width * r.height	
}

func (r *Rectangle) setWidth(width float64) {
	r.width = width
}

func (r Rectangle) area() float64{
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}


func main(){
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}

	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is ", r1.area())
	fmt.Println("Area of r2 is ", r2.area())

	fmt.Println("Area of c1 is ", c1.area())
	fmt.Println("Area of c2 is ", c2.area())

	r1.setWidth(10)
	fmt.Println("Now Area of r1 is ", r1.area())
}