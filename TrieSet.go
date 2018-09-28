package goDataStructure

type TrieSet struct {
	root *trieSetNode
	size int
}

type trieSetNode struct {
	isWord bool
	next   map[rune]*trieSetNode
}

func CreateTrieSet() *TrieSet {
	return &TrieSet{
		root: &trieSetNode{
			isWord: false,
			next:   make(map[rune]*trieSetNode),
		},
		size: 0,
	}
}

func (mt *TrieSet) GetSize() int {
	return mt.size
}

func (mt *TrieSet) Add(word Comparable) { //非递归写法 todo 完成add的递归写法
	cur := mt.root
	char := []rune(string(word.(Stringer)))

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			//newNode :=createTrieNode()
			//fmt.Println(newNode)
			cur.next[value] = &trieSetNode{
				isWord: false,
				next:   map[rune]*trieSetNode{},
			}
		}
		cur = cur.next[value]
	}
	if !cur.isWord {
		cur.isWord = true
		mt.size++
	}
}
func (mt *TrieSet) Contains(word Comparable) bool {
	cur := mt.root
	char := []rune(string(word.(Stringer)))

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			return false
		}
		cur = cur.next[value]
	}
	return cur.isWord
}
func (mt *TrieSet) IsPrefix(prefix Comparable) bool {
	cur := mt.root
	char := []rune(string(prefix.(Stringer)))

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			return false
		}
		cur = cur.next[value]
	}
	return true
}

func (mt *TrieSet) Remove(value Comparable) {
	panic("implement me")
}

func (mt *TrieSet) IsEmpty() bool {
	return mt.size == 0
}

func (mt *TrieSet) String() string {
	panic("implement me")
}
