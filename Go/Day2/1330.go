package main

import "fmt"

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	t := h*60 + m
	t -= 45

	if t < 0 {
		t += 24 * 60
	}

	fmt.Printf("%d %d", t/60, t%60)
}
