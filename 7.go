package main

import (
	"fmt"
	"sync"
)

type CustomMap struct {
	myMap map[int]string
	mutex sync.Mutex
}

func (m *CustomMap) Init() {
	m.myMap = make(map[int]string)
}

func (m *CustomMap) Set(key int, value string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.myMap[key] = value
}

func (m *CustomMap) Get(key int) string {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.myMap[key]
}

func main() {
	myMap := CustomMap{}
	myMap.Init()
	myMap.Set(1, "A")
	fmt.Println(myMap.Get(1))
}

//p.s: один в один как кэш в L0
