package main

import "fmt"

func compare(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	res := compare([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(res)
}
