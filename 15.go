/*
var justString string
func someFunc() {
	v := createHugeString(1 << 10) //создаём большую строку
	justString = v[:100]
	//инициализируем строку первыми 100 байтами строки v, и т.к строки immutable
	//то justString так и продолжит ссылаться на первые 100 байт, а "хвост" будет оставаться в памяти, и
	//не удалится сборщиком мусора
}
func main() {
	someFunc()
}
*/

// корректно:
package main

import "fmt"

var justString []byte

func createHugeString(len int) string {
	return "1234567890qwertyuiop[]"
}

// для примера взяты другие значения
func someFunc() {
	v := createHugeString(1 << 10)
	justString = make([]byte, 9)
	//Копируем нужные нам байты из v
	copy(justString, []byte(v))
}

func main() {
	someFunc()
	fmt.Println(string(justString))
}
