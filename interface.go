package data

type Queue interface {
	//队列
	DeQueue() interface{}
	EnQueue(e interface{})
	GetFront() interface{}
	GetSize() int
	IsEmpty() bool
	String() string
}
type List interface {
	//链表
	RemoveFirst() interface{}
	AddHead(e interface{})
	GetSize() int
	IsEmpty() bool
	Contains(e interface{}) bool
	String() string
}

type Stack interface {
	//栈
	GetSize() int
	IsEmpty() bool
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
	String() string
}

//树
type Tree interface {
	GetSize() int
	IsEmpty() bool
	Add(value Comparable)
	Remove(value Comparable)
	Contains(value Comparable) bool
	Min() Comparable
	Max() Comparable
	String() string
}

//Comparable 可比较的
type Comparable interface {
	Compare(c2 Comparable) int
}

//MergerAble 可合并的
type MergerAble interface {
	Merge(m2 MergerAble) MergerAble
}

//Set 集合
type Set interface {
	GetSize() int
	IsEmpty() bool
	Add(value Comparable)
	Remove(value Comparable)
	Contains(value Comparable) bool
	String() string
}

//Map 映射
type Map interface {
	Add(k interface{}, v interface{})
	Remove(k interface{}) interface{}
	Contains(key interface{}) bool
	Get(key interface{}) interface{}
	Set(key interface{}, newValue interface{})
	GetSize() int
	IsEmpty() bool
	String() string
}

//Heap 堆
type Heap interface {
	Add(e Comparable)
	Remove() Comparable
	shiftUp(index int)
	shiftDown(index int)
	GetSize() int
	IsEmpty() bool
	String() string
}

//SegmentTree 线段树
type SegmentTree interface {
	Query(l, r int) MergerAble
	Set(index int, e MergerAble)
	GetSize() int
	String() string
}

//UnionFind 并查集
type UnionFind interface {
	Union(p, q int)
	IsConnected(p, q int) bool
	GetSize() int
	String() string
}
