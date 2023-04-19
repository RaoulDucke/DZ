package main

import (
	"fmt"
	"strings"
)

func coincide(a string, b string) int {
	a = strings.ToLower(a)
	a = strings.Trim(a, "?")
	b = strings.ToLower(b)
	x := strings.Split(a, " ")
	for i, v := range x {
		if v == b {
			return i
		}
	}
	return -1

}

func main() {
	res := coincide("Ты шо лох?", "лох")
	fmt.Print(res, "\n")
}
