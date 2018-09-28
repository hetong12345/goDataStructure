package goDataStructure

type TrieMap struct {
	root *trieMapNode
	size int
}
type trieMapNode struct {
	isWord bool
	value  interface{}
	next   map[rune]*trieMapNode
}

func (tm *TrieMap) Add(k interface{}, v interface{}) { //todo 完成TrieMap 实现词频统计
	panic("implement me")
}

func (tm *TrieMap) Remove(k interface{}) interface{} {
	panic("implement me")
}

func (tm *TrieMap) Contains(key interface{}) bool {
	panic("implement me")
}

func (tm *TrieMap) Get(key interface{}) interface{} {
	panic("implement me")
}

func (tm *TrieMap) Set(key interface{}, newValue interface{}) {
	panic("implement me")
}

func (tm *TrieMap) GetSize() int {
	panic("implement me")
}

func (tm *TrieMap) IsEmpty() bool {
	panic("implement me")
}

func (tm *TrieMap) String() string {
	panic("implement me")
}
