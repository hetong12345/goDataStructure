package data

type avlTreeSet struct {
	atm *AvlTreeMap
}

func CreateAvlTreeSet() *avlTreeSet {
	return &avlTreeSet{
		atm: CreateAvlTreeMap(),
	}
}
func (ats *avlTreeSet) GetSize() int {
	return ats.atm.GetSize()
}

func (ats *avlTreeSet) IsEmpty() bool {
	return ats.atm.IsEmpty()
}

func (ats *avlTreeSet) Add(value Comparable) {
	ats.atm.Add(value, nil)
}

func (ats *avlTreeSet) Remove(value Comparable) {
	ats.atm.Remove(value)
}

func (ats *avlTreeSet) Contains(value Comparable) bool {
	return ats.atm.Contains(value)
}

func (ats *avlTreeSet) String() string {
	return ats.atm.String()
}
