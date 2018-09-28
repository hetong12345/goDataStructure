package goDataStructure

import "fmt"

type LoopQueue struct {
	data        []interface{}
	tail, front int
}

func CreateLoopQueue(cap int) *LoopQueue {
	return &LoopQueue{
		make([]interface{}, cap+1),
		0,
		0,
	}
}

func (q *LoopQueue) EnQueue(e interface{}) {
	if q.IsFull() {
		q.resize(q.GetCap() * 2)
	}
	q.data[q.tail] = e
	q.tail = (q.tail + 1) % len(q.data)
}

func (q *LoopQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		fmt.Println("队列为空！", q)
	}
	ret := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % len(q.data)
	//fmt.Println(q.GetSize(),"`````````",q.GetCap())
	if (q.GetSize() == q.GetCap()/4) && q.GetCap()/2 != 0 {
		*q = q.resize(q.GetCap() / 2)
	}
	return ret
}

func (q *LoopQueue) GetFront() interface{} {
	if q.IsEmpty() {
		fmt.Println("队列为空！")
	}
	return q.data[q.front]
}

func (q *LoopQueue) resize(newCap int) LoopQueue {
	//fmt.Println("---------do resize ---------")
	nq := CreateLoopQueue(newCap)
	*nq = nq.copyQueue(q)
	*q = *nq
	return *q
}

//func (nq *LoopQueue) copyQueue(q *LoopQueue) LoopQueue {
//	for i := 0; i < q.GetSize(); i++ {
//		nq.data[i] = q.data[(i+q.front)%len(q.data)]
//	}
//	nq.front, nq.tail = 0, q.GetSize()
//	return *nq
//}
func (q *LoopQueue) copyQueue(nq *LoopQueue) LoopQueue {
	for i := 0; i < nq.GetSize(); i++ {
		q.data[i] = nq.data[(i+nq.front)%len(nq.data)]
	}
	q.front, q.tail = 0, nq.GetSize()
	return *q
}

func (q *LoopQueue) GetCap() int {
	return len(q.data) - 1
}

func (q *LoopQueue) GetSize() int {
	return (q.tail - q.front + len(q.data)) % len(q.data)
}

func (q *LoopQueue) IsEmpty() bool {
	return q.front == q.tail
}
func (q *LoopQueue) IsFull() bool {
	return (q.tail+1)%len(q.data) == q.front
}
func (q *LoopQueue) String() string {
	str := fmt.Sprint("Queue:")
	str += fmt.Sprint(("front ["))
	for i := q.front; i != q.tail; i = (i + 1) % len(q.data) {
		if q.data[i] == nil {
			continue
		}
		str += fmt.Sprint((*q).data[i])
		if (i+1)%len(q.data) != q.tail {
			str += fmt.Sprint(",")
		}
	}
	str += fmt.Sprint("] tail")
	return str
}
