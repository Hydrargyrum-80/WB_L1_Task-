package main

import "fmt"

// "Родительская" структура
type Human struct{}

func (H *Human) foo1() {
	fmt.Println("Human:foo1")
}

func (H *Human) foo2() {
	fmt.Println("Human:foo2")
}

// "Дочерняя" структура
type Action struct {
	//Т.к Go не реализует наследование в привычном его понимании (как в C++, C#, Java и др.)используется встраивание (композиция)
	Human //Предпочтительный вариант. Возможно явное переопределение (пример ниже), обращение к методам "базового" класса напрямую из дочернего: Action.foo2() (вывод: Human:foo2)

	//a Human //Не предпочтительный вариант. Невозможно явное переопределение методов (но возможно существование методов с одинаковыми именами,
	//т.к так же невозможен "прямой" вызов метода базовой структуры из дочерней структуры: Action.foo2() - метод отсутствует; Action.a.foo2() (вывод: Human:foo2)
}

// "Перекрываем метод foo1() базовой структуры
func (A *Action) foo1() {
	fmt.Println("Action:foo1")
}

func main() {
	human := Human{}
	//Вызов методов human
	human.foo1()
	human.foo2()

	action := Action{}
	//Вызов собственного метода Action
	action.foo1()
	//Вызов метода базового human
	action.Human.foo1()
	//Вызов "унаследованного" метода  human
	action.foo2()
}
