package main

import (
	"fmt"
)

func main() {
	var val int64 = 0
	n := 5
	v := int64(1 << n) //сдвиг на n бит
	val = val ^ v      // 0 xor 1 = 1 и соответственно 1 xor 1 = 0
	fmt.Println("Result val: ", val)
}
