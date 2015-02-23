package solutions

import "fmt"

/** The sum of the squares of the first ten natural numbers is,
 *	12 + 22 + ... + 102 = 385
 *	The square of the sum of the first ten natural numbers is,
 *
 *	(1 + 2 + ... + 10)2 = 552 = 3025
 *	Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
 *
 *	Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
 */
func SumSquareDifference() {
	max := 100
	s1 := 0

	for i := 0; i <= max; i = i + 1 {
		s1 = s1 + i*i
	}

	fmt.Println("s1 = ", s1)

	s2 := 0

	for i := 0; i <= max; i = i + 1 {
		s2 = s2 + i
	}

	s2 = s2 * s2
	fmt.Println("s2 = ", s2)

	fmt.Println("result = ", s2-s1)

}
