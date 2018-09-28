package goDataStructure

import "sort"

type Student struct {
	Name  string
	Score int
}

func (p *Student) Compare(c2 Comparable) int {
	return p.Score - c2.(*Student).Score
}

type Integer int

func (i Integer) Merge(i2 MergerAble) MergerAble {
	return Integer(int(i) + int(i2.(Integer)))
}

func (i Integer) Compare(c2 Comparable) int {
	return int(i) - int(c2.(Integer))
}

type Stringer string

func (s Stringer) Compare(s2 Comparable) int {
	if s == s2 {
		return 0
	}
	if sort.StringsAreSorted([]string{string(s), string(s2.(Stringer))}) {
		return -1
	}
	return 1
}
