package solutions

// the smallest positive number that is evenly divisible by all of the numbers from min to max
func SmallestMultiple(min int, max int) int {
	flag := true
	i := max
	for ; flag; i = i + 1 {
		flag = dividable(i, min, max)
	}
	return i - 1
}

func dividable(d int, min int, max int) bool {
	for j := min; j <= max; j = j + 1 {
		if d%j != 0 {
			return true
		}
	}
	return false
}
