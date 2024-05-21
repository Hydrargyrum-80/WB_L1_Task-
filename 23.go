package main

import "fmt"

func main() {
	i := 5
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println(slice)
}
