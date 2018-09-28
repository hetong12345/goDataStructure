package data

import (
	"fmt"
)

type NodeList struct {
	dummyHead *node
	size      int
}
type node struct {
	value interface{}
	next  *node
}

func CreateNodeList() *NodeList {
	return &NodeList{
		dummyHead: &node{
			value: nil,
			next:  nil,
		},
		size: 0,
	}
}

func (nl *NodeList) GetSize() int {
	return nl.size
}

func (nl *NodeList) IsEmpty() bool {
	return nl.size == 0
}

func (nl *NodeList) AddHead(e interface{}) {
	nl.dummyHead.next = &node{
		value: e,
		next:  nl.dummyHead.next,
	}
	nl.size++
}

func (nl *NodeList) AddLast(e interface{}) {
	nl.AddIndex(nl.size, e)
}

//func (nl *NodeList) AddIndex(index int, e interface{}) {//循环实现
//	if index < 0 || index > nl.size {
//		panic("fei fa wei zhi")
//	}
//	prev := &nl.dummyHead
//	for i := 0; i < index; i++ {
//		prev = prev.next
//	}
//	newNode := node{
//		value: e,
//		next:  prev.next,
//	}
//	prev.next = &newNode
//	nl.size++
//}
func (nl *NodeList) AddIndex(index int, e interface{}) { //递归实现
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	addIndex(nl.dummyHead, index, e)
	nl.size++
}
func addIndex(prev *node, index int, e interface{}) {
	if index == 0 {
		prev.next = &node{e, prev.next}
		return
	}
	addIndex(prev.next, index-1, e)
}

//func (nl *NodeList) Get(index int) interface{} {//循环实现
//	if index < 0 || index > nl.size {
//		panic("fei fa wei zhi")
//	}
//	current := nl.dummyHead.next
//	for i := 0; i < index; i++ {
//		current = current.next
//	}
//	return current.value
//}
func (nl *NodeList) Get(index int) interface{} { //递归实现
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	return get(nl.dummyHead.next, index)
}
func get(cur *node, index int) interface{} {
	if index == 0 {
		return cur.value
	}
	return get(cur.next, index-1)
}
func (nl *NodeList) GetFirst() interface{} {
	return nl.Get(0)
}
func (nl *NodeList) GetLast() interface{} {
	return nl.Get(nl.size - 1)
}

//func (nl *NodeList) Set(index int, e interface{}) {//循环实现
//	if index < 0 || index > nl.size {
//		panic("fei fa wei zhi")
//	}
//	current := nl.dummyHead.next
//	for i := 0; i < index; i++ {
//		current = current.next
//	}
//	current.value = e
//}
func (nl *NodeList) Set(index int, e interface{}) {
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	set(nl.dummyHead.next, index, e)
}
func set(cur *node, index int, e interface{}) {
	if index == 0 {
		cur.value = e
		return
	}
	set(cur.next, index-1, e)
}
func (nl *NodeList) SetFirst(e interface{}) {
	nl.Set(0, e)
}
func (nl *NodeList) SetLast(e interface{}) {
	nl.Set(nl.size-1, e)
}
func (nl *NodeList) Contains(e interface{}) bool {
	current := nl.dummyHead.next
	for current.value != nil {
		if current.value == e {
			return true //递归实现
		}
		current = current.next
	}
	return false
}

//func (nl *NodeList) Remove(index int) interface{} {
//	if index < 0 || index > nl.size {
//		panic("fei fa wei zhi")
//	}
//	prev := &nl.dummyHead
//	for i := 0; i < index; i++ {
//		prev = prev.next
//	}
//	ret := prev.next
//	prev.next = ret.next
//	ret.next = nil
//
//	nl.size--
//	return ret.value
//}
func (nl *NodeList) Remove(index int) interface{} {
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	nl.size--
	return removeNode(nl.dummyHead, index)
}
func removeNode(prev *node, index int) interface{} {
	if index == 0 {
		ret := prev.next
		prev.next = ret.next
		ret.next = nil
		return ret.value
	}
	return removeNode(prev.next, index-1)
}
func (nl *NodeList) RemoveFirst() interface{} {
	return nl.Remove(0)
}
func (nl *NodeList) RemoveLast() interface{} {
	return nl.Remove(nl.size - 1)
}
func (nl *NodeList) String() string {
	str := fmt.Sprint("List:")
	str += fmt.Sprint("Head ")
	current := nl.dummyHead.next
	for {
		str += fmt.Sprint(current.value)
		if current.next != nil {
			str += fmt.Sprint("->")
			current = current.next
		} else {
			break
		}
	}
	str += fmt.Sprint("->nil")
	return str
}
