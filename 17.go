package main

import "fmt"

func binarySearch(arr []int, findEl int) int {
	var (
		left   int = 0
		right  int = len(arr) - 1
		result int = 0
	)
	for left <= right {
		result = (left + right) / 2
		if findEl >= arr[result] && findEl <= arr[result] {
			return result
		} else if findEl < arr[result] {
			right = result - 1
		} else {
			left = result + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(binarySearch(arr, 10))
}
