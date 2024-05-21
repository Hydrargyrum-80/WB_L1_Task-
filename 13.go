package main

import "fmt"

func main() {
	var (
		a int = 0
		b int = 1
	)
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
