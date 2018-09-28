package data

type BSTreeSet struct {
	bst *BSTreeMap
}

func CreateBSTreeSet() *BSTreeSet {
	return &BSTreeSet{CreateBSTreeMap()}
}

func (ts *BSTreeSet) GetSize() int {
	return ts.bst.GetSize()
}

func (ts *BSTreeSet) IsEmpty() bool {
	return ts.bst.IsEmpty()
}

func (ts *BSTreeSet) Add(value Comparable) {
	ts.bst.Add(value, nil)
}

func (ts *BSTreeSet) Remove(value Comparable) {
	ts.bst.Remove(value)
}

func (ts *BSTreeSet) Contains(value Comparable) bool {
	return ts.bst.Contains(value)
}
func (ts *BSTreeSet) String() string {
	return ts.bst.String()
}
