package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	resultSet := map[string]struct{}{}
	for _, i := range sequence {
		resultSet[i] = struct{}{}
	}
	fmt.Println(resultSet)
}
