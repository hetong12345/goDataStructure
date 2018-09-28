package goDataStructure

import "fmt"

type ListMap struct {
	dummyHead *mapNode
	size      int
}

type mapNode struct {
	key   interface{}
	value interface{}
	next  *mapNode
}

func CreateListMap() *ListMap {
	return &ListMap{
		dummyHead: &mapNode{
			key:   nil,
			value: nil,
			next:  nil,
		},
		size: 0,
	}
}

func (lm *ListMap) getMapNode(key interface{}) *mapNode {
	cur := lm.dummyHead.next
	for cur != nil {
		if cur.key == key {
			return cur
		}
		cur = cur.next
	}
	return nil
}

func (lm *ListMap) Add(k interface{}, v interface{}) {
	node := lm.getMapNode(k)
	if node == nil {
		lm.dummyHead.next = &mapNode{
			key:   k,
			value: v,
			next:  lm.dummyHead.next,
		}
		lm.size++
	} else {
		node.value = v
	}
}

func (lm *ListMap) Remove(k interface{}) interface{} {
	prev := lm.dummyHead
	for prev != nil {
		if prev.next.key == k {
			break
		}
		prev = prev.next
	}
	if prev != nil {
		del := prev.next
		prev.next = del.next
		del.next = nil
		lm.size--
		return del.value
	}
	return nil
}

func (lm *ListMap) Contains(key interface{}) bool {
	return lm.getMapNode(key) == nil
}

func (lm *ListMap) Get(key interface{}) interface{} {
	node := lm.getMapNode(key)
	if node != nil {
		return node.value
	}
	return nil
}

func (lm *ListMap) Set(key interface{}, newValue interface{}) {
	node := lm.getMapNode(key)
	if node == nil {
		panic("bu cun zai zhe yang de key")
	} else {
		node.value = newValue
	}
}

func (lm *ListMap) GetSize() int {
	return lm.size
}

func (lm *ListMap) IsEmpty() bool {
	return lm.size == 0
}
func (lm *ListMap) String() string {
	str := fmt.Sprint("List:")
	str += fmt.Sprint("Head ")
	current := lm.dummyHead.next
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
