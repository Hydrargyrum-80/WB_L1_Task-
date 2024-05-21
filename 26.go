package main

import (
	"fmt"
	"strings"
)

func uniqueCharCheck(str string) bool {
	str = strings.ToLower(str) //приводим строку к нижнему регистру
	uniqueMap := make(map[string]struct{})
	for _, v := range str { //проходим по всей строке
		if _, ok := uniqueMap[string(v)]; !ok { //если элемента нет в мапе
			uniqueMap[string(v)] = struct{}{} //добавляем
		} else {
			return false //если есть - символ в строке не уникален
		}
	}
	return true
}

func main() {

	str1 := "12345QWERTYUIOP"
	str2 := "12345QWERTYUIOPQ"
	str3 := "12345QWERTYUIOPq"
	str4 := "1234567890qwertyuiOP[]ASDF"
	fmt.Println(str1, " ", uniqueCharCheck(str1))
	fmt.Println(str2, " ", uniqueCharCheck(str2))
	fmt.Println(str3, " ", uniqueCharCheck(str3))
	fmt.Println(str4, " ", uniqueCharCheck(str4))
}
