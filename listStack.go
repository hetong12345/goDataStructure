package goDataStructure

import "fmt"

type ListStack struct {
	nl *NodeList
}

func CreateListStack() *ListStack {
	return &ListStack{CreateNodeList()}
}
func (ls *ListStack) GetSize() int {
	return ls.nl.GetSize()
}

func (ls *ListStack) IsEmpty() bool {
	return ls.nl.IsEmpty()
}

func (ls *ListStack) Push(e interface{}) {
	ls.nl.AddHead(e)
}

func (ls *ListStack) Pop() interface{} {
	return ls.nl.RemoveFirst()
}

func (ls *ListStack) Peek() interface{} {
	return ls.nl.GetFirst()
}

func (ls *ListStack) String() string {
	str := fmt.Sprint("Stack:")
	str += fmt.Sprint("top [")

	str += fmt.Sprint(ls.nl.String())

	str += fmt.Sprint("]")
	//fmt.Println(str)
	return str
}
