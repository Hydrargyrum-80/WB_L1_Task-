package main

import (
	"fmt"
	"reflect"
)

func printType(T interface{}) { //T any аналогично
	//можно было бы использовать any.(type), однако тогда невозможно определить "абстрактный" канал
	//без конкретного типа его значения
	tVal := reflect.ValueOf(T)
	switch tVal.Kind() {
	case reflect.Int:
		{
			fmt.Println("int")
		}
	case reflect.String:
		{
			fmt.Println("string")
		}
	case reflect.Bool:
		{
			fmt.Println("bool")
		}
	case reflect.Chan:
		{
			fmt.Println("chan")
		}
	default:
		fmt.Println("unhandled default case")
	}
}

func main() {
	printType(1)
	printType("1")
	printType(false)
	printType(make(chan int))
}
