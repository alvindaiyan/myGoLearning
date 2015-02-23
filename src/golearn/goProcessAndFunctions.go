package main

import (
	"fmt"
	"os"
)

func main() {

	// try defer
	try_defer()

	// try function type
	integers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println("the odd numbers are", filter(integers, isOdd))
	fmt.Println("the even numbers are", filter(integers, isEven))

	// panic and recover
	foo()
	fmt.Println("returned normally from f")

}

// defer
func try_defer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

// try function type
type testInt func(int) bool // declare a function type

func isOdd(x int) bool {
	return x%2 != 0
}

func isEven(x int) bool {
	return x%2 == 0
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

// try Panic and Recover
var user = os.Getenv("USER")

func f() {
	if user == "" {
		panic("no value for $USER")
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("recover f")
			b = true
		}
	}()
	f()
	return
}

func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in foo", r)
		}
	}()
	fmt.Println("calling g")
	g(0)
	fmt.Println("returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
