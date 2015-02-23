package solutions

import (
	"fmt"
	"math"
)

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

func SpecialPythagoreaTriplet() {
	var max float64 = 500 //333
	var a float64 = 1
	for ; a <= max; a = a + 1 {
		var b float64 = 1
		for ; b < max; b++ {
			c := math.Sqrt(a*a + b*b)
			if a+b+c == 1000 {
				fmt.Println(a, " &", b, " & ", c)
				fmt.Println("Product is ", a*b*c)
			}
		}

	}
}
