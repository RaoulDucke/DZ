package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "sdfsfgvse4&? w45 wq43trq3"
	a = regexp.MustCompile(`\D+`).ReplaceAllString(a, "")
	fmt.Println(a)
}
