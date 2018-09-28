package goDataStructure

import "fmt"

type QuickFindUnionFind struct {
	id []int
}

func CreateQuickFindUnionFind(size int) *QuickFindUnionFind {
	arr := make([]int, size)
	for key := range arr {
		arr[key] = key
	}
	return &QuickFindUnionFind{
		id: arr,
	}
}

func (qf *QuickFindUnionFind) find(x int) int {
	if x < 0 || x >= qf.GetSize() {
		panic("x is out of bound")
	}
	return qf.id[x]
}

func (qf *QuickFindUnionFind) Union(p, q int) {
	pId := qf.find(p)
	qId := qf.find(q)

	if pId == qId {
		return
	}
	//fmt.Println(pId,qId)
	for key, value := range qf.id {
		if value == pId {
			qf.id[key] = qId
		}
	}
}

func (qf *QuickFindUnionFind) IsConnected(p, q int) bool {
	return qf.find(p) == qf.find(q)
}

func (qf *QuickFindUnionFind) GetSize() int {
	return len(qf.id)
}

func (qf *QuickFindUnionFind) String() string {
	return fmt.Sprintln(qf.id)
}
