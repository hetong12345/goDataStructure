package goDataStructure

type PriorityQueue struct {
	mh *MaxHeap
}

func CreatePriorityQueue() *PriorityQueue {

	return &PriorityQueue{
		CreateMaxHeap(),
	}
}

func (pq *PriorityQueue) DeQueue() interface{} {
	return pq.mh.Remove()
}

func (pq *PriorityQueue) EnQueue(e interface{}) {
	ce, ok := e.(Comparable)
	if !ok {
		panic("e is no a comparable ver")
	}
	pq.mh.Add(ce)
}

func (pq *PriorityQueue) GetFront() interface{} {
	return pq.mh.getMax()
}

func (pq *PriorityQueue) GetSize() int {
	return pq.mh.GetSize()
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.mh.IsEmpty()
}

func (pq *PriorityQueue) String() string {
	return pq.mh.String()
}
