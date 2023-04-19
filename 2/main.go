package main

import "fmt"

func reverse(a []int) []int {
	for b, c := 0, len(a)-1; b < c; b, c = b+1, c-1 {
		a[b], a[c] = a[c], a[b]
	}
	return a
}

func main() {
	res := []int{3, 2, 1}
	reverse(res)
	fmt.Println(res)
}
