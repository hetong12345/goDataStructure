package goDataStructure

import (
	"fmt"
	"log"
)

//重写动态数组
type arr struct {
	data []interface{}
	size int
}

//func CreateArray(cap int) *arr {
//	a := []interface{}{}
//	for i := 0; i < cap; i++ {
//		a = append(a, 0)
//	}
//	return &arr{
//		data: a,
//		size: 0,
//	}
//}
func CreateArray(cap int) *arr {
	return &arr{
		data: make([]interface{}, cap),
		size: 0,
	}
}
func (a *arr) GetSize() int {
	return a.size
}
func (a *arr) GetCap() int {
	return len(a.data)
}
func (a *arr) AddLast(e interface{}) {
	a.Add(a.size, e)
	return
}
func (a *arr) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("shu zu qu chu chu cuo")
	}
	return a.data[index]
}
func (a *arr) GetFirst() interface{} {
	return a.Get(0)
}

func (a *arr) GetLast() interface{} {
	return a.Get(a.size - 1)
}
func (a *arr) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("shu zu xiu gai chu cuo")
	}
	a.data[index] = e
}
func (a *arr) AddFirst(e interface{}) {
	a.Add(0, e)
	return
}
func (a *arr) Add(index int, e interface{}) {

	if index < 0 || index > a.size {
		log.Println("fei fa weizhi ")
		return
	}
	if a.size == len(a.data) && len(a.data)/2 != 0 {
		*a = a.resize(a.size * 2)
		//fmt.Println(*a)
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
	return

}

func (a *arr) resize(len int) arr {
	na := CreateArray(len)
	*na = na.copyArr(a)
	*a = *na
	return *a
}

//func (na *arr) copyArr(a *arr) arr {
//	for i := 0; i < a.size; i++ {
//		na.data[i] = a.data[i]
//	}
//	na.size = a.size
//	return *na
//}
func (a *arr) copyArr(na *arr) arr {
	for i := 0; i < na.size; i++ {
		a.data[i] = na.data[i]
	}
	a.size = na.size
	return *a
}
func (a *arr) Remove(index int) interface{} {
	if index < 0 || index > a.size {
		panic("fei fa weizhi")
	}
	ret := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--

	if a.size == len(a.data)/4 {
		*a = (a.resize(len(a.data) / 2))
	}
	return ret

}
func (a *arr) RemoveFirst() interface{} {
	return a.Remove(0)
}
func (a *arr) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}
func (a *arr) RemoveElement(e interface{}) {
	index := a.FindFirst(e)
	if index != -1 {
		a.Remove(index)
	}
}
func (a *arr) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return true
		}
	}
	return false
}
func (a *arr) FindFirst(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}
func (a *arr) IsEmpty() bool {
	return a.size == 0
}

func (a *arr) String() string {
	str := fmt.Sprintf("Array:size=%d,cap=%d\n", a.size, len(a.data))
	str += fmt.Sprint(("["))
	for i := 0; i < a.size; i++ {
		str += fmt.Sprint((*a).data[i])
		if i != a.size-1 {
			str += fmt.Sprint(",")
		}
	}
	str += fmt.Sprint("]")
	//fmt.Println(str)
	return str
}
