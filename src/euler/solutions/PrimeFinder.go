package solutions

/**
 * find a given numbers prime number start from 2
 */
func PrimeFind(length int) []int {
	prime := make([]int, length)
	count := 0
	for i := 2; count < length; i = i + 1 {
		if isPrime(i, prime) {
			prime[count] = i
			count = count + 1
		}
	}
	return prime
}

func isPrime(i int, primes []int) bool {
	for j := range primes {
		if primes[j] != 0 && i%primes[j] == 0 {
			return false
		}
	}
	return true

}
