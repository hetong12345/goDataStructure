package goDataStructure

import "fmt"

type MaxHeap struct {
	array []Comparable
}

func CreateMaxHeap() *MaxHeap {
	return &MaxHeap{
		array: []Comparable{},
	}
}

func (mh *MaxHeap) parent(index int) int {
	if index == 0 {
		panic("index haven't parent ")
	}
	return (index - 1) / 2
}

func (mh *MaxHeap) leftChild(index int) int {
	return index*2 + 1
}

func (mh *MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

func (mh *MaxHeap) Add(e Comparable) {
	mh.array = append(mh.array, e)
	mh.shiftUp(mh.GetSize() - 1)
}

func (mh *MaxHeap) shiftUp(index int) {
	for index > 0 && mh.array[mh.parent(index)].Compare(mh.array[index]) < 0 {
		mh.array[mh.parent(index)], mh.array[index] = mh.array[index], mh.array[mh.parent(index)]
		index = mh.parent(index)
	}
}

func (mh *MaxHeap) Remove() Comparable {
	ret := mh.getMax()
	mh.array[0], mh.array[mh.GetSize()-1] = mh.array[mh.GetSize()-1], mh.array[0]
	mh.array = mh.array[0 : len(mh.array)-1]
	mh.shiftDown(0)
	return ret
}

func (mh *MaxHeap) shiftDown(index int) {
	for mh.exist(mh.leftChild(index)) {
		max := mh.leftChild(index)

		if mh.exist(mh.rightChild(index)) && mh.array[mh.rightChild(index)].Compare(mh.array[mh.leftChild(index)]) > 0 {
			max++
		}
		if mh.array[index].Compare(mh.array[max]) >= 0 {
			return
		}
		mh.array[max], mh.array[index] = mh.array[index], mh.array[max]
		index = max
	}
}
func (mh *MaxHeap) Replace(newEle Comparable) Comparable {
	ret := mh.array[0]
	mh.array[0] = newEle
	mh.shiftDown(0)
	return ret
}
func Heapify(arr []Comparable) *MaxHeap {
	ret := &MaxHeap{
		array: arr,
	}
	for i := ret.parent(len(arr) - 1); i >= 0; i-- {
		ret.shiftDown(i)
	}
	return ret
}
func (mh *MaxHeap) exist(index int) bool {
	return index < len(mh.array)
}
func (mh *MaxHeap) GetSize() int {
	return len(mh.array)
}

func (mh *MaxHeap) IsEmpty() bool {
	return len(mh.array) == 0
}

func (mh *MaxHeap) String() string {
	return fmt.Sprint(mh.array)
}
func (mh *MaxHeap) getMax() Comparable {
	if mh.IsEmpty() {
		panic("heap is empty")
	}
	return mh.array[0]
}
