package main

import (
	"fmt"
	"strings"
)

type arr struct {
	strs []string
}

func (a *arr) Add(str string) {
	a.strs = append(a.strs, str)
}

func main() {
	var arrObj arr

	arrObj.Add("str0")
	arrObj.Add("str1")
	arrObj.Add("str2")
	arrObj.Add("str3")
	arrObj.Add("str4")
	arrObj.Add("str5")
	arrObj.Add("str6")

	for _, str := range arrObj.strs {
		fmt.Println("... ", str)
	}

	fmt.Printf("[%q]", strings.Trim(" !!! Achtung! Achtung! !!! ", " "))
}
