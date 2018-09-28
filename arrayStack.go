package goDataStructure

import "fmt"

type ArrayStack struct {
	a *arr
}

func CreateArrayStack(cap int) *ArrayStack {
	return &ArrayStack{CreateArray(cap)}
}

func (as *ArrayStack) GetSize() int {
	return as.a.GetSize()
}

func (as *ArrayStack) IsEmpty() bool {
	return as.a.IsEmpty()
}

func (as *ArrayStack) Push(e interface{}) {
	as.a.AddLast(e)
}

func (as *ArrayStack) Pop() interface{} {
	return as.a.RemoveLast()
}

func (as *ArrayStack) Peek() interface{} {
	return as.a.GetLast()
}
func (as *ArrayStack) String() string {
	str := fmt.Sprint("Stack:")
	str += fmt.Sprint(("["))
	for i := 0; i < as.a.size; i++ {
		str += fmt.Sprint((*as).a.data[i])
		if i != as.a.size-1 {
			str += fmt.Sprint(",")
		}
	}
	str += fmt.Sprint("] top")
	//fmt.Println(str)
	return str
}
