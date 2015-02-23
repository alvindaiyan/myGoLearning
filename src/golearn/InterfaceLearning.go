package main

import "fmt"

type Human struct{
	name string
	age int
	phone string
}

type Student struct{
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human) SayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string){
	fmt.Println("sing sing sing", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s \n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}



func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "4444-222-xxx"}, "Golan Inc.", 1000}    
	tom := Employee{Human{"Tom", 37, "2222-333-444"}, "google Inc.", 1000}


	var i Men

	i = mike
	fmt.Println("this is Mike, a Student:")
	i.SayHi()
	i.Sing("Nov Nov")

	i = sam
	fmt.Println("this is Sam, a Employee:")
	i.SayHi()
	i.Sing("Gal Gal")

	i = paul 
	fmt.Println("this is Paul, a Employee:")
	i.SayHi()
	i.Sing("Foo Foo")

	i = tom 
	fmt.Println("this is Tom, a Employee:")
	i.SayHi()
	i.Sing("Dec Dec")


	x := make([]Men, 3)
	x[0], x[1], x[2] = sam, tom, mike
	for _, value := range x {
		value.SayHi()
	}
}