package main

import "fmt"

func main() {
	var (
		a int64 = 2097152
		b int64 = 2097152
	)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a + b)
	fmt.Println(a - b)
}
