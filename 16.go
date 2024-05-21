package main

import "fmt"

func qSort(a []int) {
	if len(a) <= 1 {
		return
	}
	left, right := 0, len(a)-1
	p := len(a) / 2
	a[p], a[right] = a[right], a[p]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	qSort(a[:left])
	qSort(a[left+1:])
}

func main() {
	arr := []int{-10, 9, 8, 7, 6, 100534234, 5, 4, 3, 2, 1}
	qSort(arr)
	fmt.Println(arr)
}
