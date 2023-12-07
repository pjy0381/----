package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result := "=="

	if a > b {
		result = ">"
	} else if a < b {
		result = "<"
	}

	fmt.Println(result)
}
