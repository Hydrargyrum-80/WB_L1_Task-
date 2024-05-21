package main

import "fmt"

type IArray interface {
	moveAndExpand()
}

// конкретный массив
type Array struct {
}

func (a *Array) moveAndExpand() {
	fmt.Println("Array: moveAndExpand()")
}

type IList interface {
	insert()
}

// конкретный список
type List struct {
}

func (l *List) insert() {
	fmt.Println("List: insert()")
}

// адаптер
type CListAdapter struct {
	array Array //храним в адаптере массив
}

// Теперь в с Array как бы можно применять метод insert(), хотя на самом деле вызывается moveAndExpand()
func (a *CListAdapter) insert() {
	a.array.moveAndExpand()
}

func main() {
	arr := Array{}
	arr.moveAndExpand()
	list := List{}
	list.insert()

	//"превращаем" массив в список
	var listAdapter CListAdapter = CListAdapter{arr}
	listAdapter.insert()
}
