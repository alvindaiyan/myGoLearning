package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Frame struct {
	upper int
	lower int
}

func LargestProductSeries(input string) int {
	array := initArray(input)
	frame := Frame{0, 13}
	max := 0
	for frame.lower < len(array) {
		product := 1
		for i := frame.upper; i < frame.lower; i++ {
			product = product * array[i]
		}
		if product > max {
			max = product
		}

		frame.upper++
		frame.lower++
	}

	return max
}

func initArray(input string) []int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
	count := 0
	array := make([]int, 1000)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			array[count] = i
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return array
}
