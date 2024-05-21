package main

import (
	"fmt"
	"strings"
)

func Reverse20(s string) string { //используем функцию из предыдущего задания
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseWords(str string) string {
	splitStr := strings.Split(str, " ") //делим строку на слова
	for i := range splitStr {
		splitStr[i] = Reverse20(splitStr[i])
	}
	return strings.Join(splitStr, " ") //складываем все строки вместе
}

func main() {
	str := "Hello World and Go"
	fmt.Println(str)
	fmt.Println(ReverseWords(str))
}
