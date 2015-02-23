package solutions

import (
	"fmt"
)

func LargestPrimeFactor() {
	var max int64 = 997799 // 71 839 1471 6857 998001
	// fmt.Println(max / 907)
	for j := max; j >= max-90090; j = j - 10010 {
		for i := j; i >= j-9900; i = i - 1100 {
			fmt.Println("tried ", i)
			getPrimesMax(i)
		}
	}
}

func getPrimesMax(max int64) {
	var i int64 = 100
	for ; i < int64(1000); i++ {
		// if isPrime(i, primes) {
		if i != 0 && i != 1 && max%i == 0 && max/i < 1000 && max/i > 99 {
			fmt.Println("number: ", max)
			fmt.Println(i)
		}
		// primes[i] = true
		// }
	}
}

// func isPrime(i int64, primes []bool) bool {
// 	for idx := 0; idx < len(primes); idx = idx + 1 {
// 		if idx != 0 && idx != 1 && primes[idx] && i%int64(idx) == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
