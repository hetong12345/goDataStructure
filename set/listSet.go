package data

type ListSet struct {
	il *ImproveList
}

func CreateListSet() *ListSet {
	return &ListSet{CreateImproveList()}
}
func (ls *ListSet) GetSize() int {
	return ls.il.GetSize()
}

func (ls *ListSet) IsEmpty() bool {
	return ls.il.IsEmpty()
}

func (ls *ListSet) Add(value Comparable) {
	if !ls.il.Contains(value) {
		ls.il.AddHead(value)
	}
}

func (ls *ListSet) Remove(value Comparable) {
	ls.il.RemoveFirst()
}

func (ls *ListSet) Contains(value Comparable) bool {
	return ls.il.Contains(value)
}
func (ls *ListSet) String() string {
	return ls.il.String()
}
