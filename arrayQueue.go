package goDataStructure

import "fmt"

type ArrayQueue struct {
	a *arr
}

//noinspection GoUnusedExportedFunction
func CreateArrayQueue(cap int) *ArrayQueue {
	return &ArrayQueue{CreateArray(cap)}
}

func (q *ArrayQueue) GetSize() int {
	return q.a.GetSize()
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.a.IsEmpty()
}

func (q *ArrayQueue) EnQueue(e interface{}) {
	q.a.AddLast(e)
}

func (q *ArrayQueue) DeQueue() interface{} {
	return q.a.RemoveFirst()
}

func (q *ArrayQueue) GetFront() interface{} {
	return q.a.GetFirst()
}
func (q *ArrayQueue) GetCap() int {
	return q.a.GetCap()
}
func (q *ArrayQueue) String() string {
	str := fmt.Sprint("Queue:")
	str += fmt.Sprint(("front ["))
	for i := 0; i < q.a.size; i++ {
		str += fmt.Sprint((*q).a.data[i])
		if i != q.a.size-1 {
			str += fmt.Sprint(",")
		}
	}
	str += fmt.Sprint("] tail")
	//fmt.Println(str)
	return str
}
