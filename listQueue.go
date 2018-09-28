package goDataStructure

import "fmt"

type ListQueue struct {
	il *ImproveList
}

func CreateListQueue() *ListQueue {
	return &ListQueue{
		il: CreateImproveList(),
	}
}

func (lq *ListQueue) DeQueue() interface{} {
	return lq.il.RemoveFirst()
}

func (lq *ListQueue) EnQueue(e interface{}) {
	lq.il.AddLast(e)
}

func (lq *ListQueue) GetFront() interface{} {
	return lq.il.GetFirst()
}

func (lq *ListQueue) GetSize() int {
	return lq.il.GetSize()
}

func (lq *ListQueue) IsEmpty() bool {
	return lq.il.IsEmpty()
}

func (lq *ListQueue) String() string {
	str := fmt.Sprint("Queue:")
	str += fmt.Sprint("front [")
	head := lq.il.head
	for {
		str += fmt.Sprint(head.value)
		if head.next != nil {
			str += fmt.Sprint("->")
			head = head.next
		} else {
			break
		}
	}
	str += fmt.Sprint("] tail")
	//fmt.Println(str)
	return str
}
