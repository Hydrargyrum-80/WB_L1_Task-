package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	resultMap := make(map[int][]float32)
	for _, i := range numbers {
		var index int
		/*
			т.к в задании не написано конкретно, к какому десятку должно быть объединение, из примера:
			-20:{-25.0, -27.0, -21.0} - к бОльшему, i<0;
			10:{13.0, 19.0, 15.5} - к меньшему, i>0;
			Чтобы соответствовать примеру, задано следующее условие:
			(float64(i)/10.0 - определяем десяток
			Floor - округляем в меньшую сторону; Ceil - в бОльшую
		*/
		if i > 0 {
			index = int(math.Floor(float64(i)/10.0) * 10)
		} else {
			index = int(math.Ceil(float64(i)/10.0) * 10)
		}
		if _, ok := resultMap[index]; !ok {
			resultMap[index] = []float32{}
		}
		resultMap[index] = append(resultMap[index], i)
	}
	fmt.Println(resultMap)
}
