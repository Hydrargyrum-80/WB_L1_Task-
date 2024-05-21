package main

import "fmt"

func main() {
	//создаём множества
	set1 := map[int]struct{}{}
	set2 := map[int]struct{}{}
	//заполняем
	for i := 0; i < 10; i++ {
		set1[i] = struct{}{}
	}
	for i := 8; i < 15; i++ {
		set2[i] = struct{}{}
	}
	resultSet := map[int]struct{}{}
	//проходимся по элементам первого множества
	for i := range set1 {
		if _, ok := set2[i]; ok { //если элемент из первого множества есть во втором - добавляем в результирующее
			resultSet[i] = struct{}{}
		}
	}
	fmt.Println("set1: ", set1)
	fmt.Println("set2: ", set2)
	fmt.Println("resultSet: ", resultSet)
}
