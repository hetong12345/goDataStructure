package data

import (
	"fmt"
	"math"
)

type AvlTreeMap struct {
	root *avlTreeMapNode
	size int
}

type avlTreeMapNode struct {
	key         Comparable
	value       interface{}
	left, right *avlTreeMapNode
	height      int
}

func CreateAvlTreeMap() *AvlTreeMap {
	return &AvlTreeMap{
		root: nil,
		size: 0,
	}
}

func (atm *AvlTreeMap) getHeight(node *avlTreeMapNode) int {
	if node == nil {
		return 0
	}
	return node.height
}
func (atm *AvlTreeMap) IsBST() bool {
	arr := []Comparable{}
	atm.midOrder(atm.root, arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1].Compare(arr[i]) > 0 {
			return false
		}
	}
	return true
}

func (atm *AvlTreeMap) midOrder(node *avlTreeMapNode, arr []Comparable) {
	if node == nil {
		return
	}
	atm.midOrder(node.left, arr)
	arr = append(arr, node.key)
	atm.midOrder(node.right, arr)
}
func (atm *AvlTreeMap) getBalanceFactor(node *avlTreeMapNode) int {
	if node == nil {
		return 0
	}
	return atm.getHeight(node.left) - atm.getHeight(node.right)
}
func (atm *AvlTreeMap) IsBalanced() bool {
	return atm.isBalanced(atm.root)
}
func (atm *AvlTreeMap) isBalanced(node *avlTreeMapNode) bool {
	if node == nil {
		return true
	}

	balanceFactor := atm.getBalanceFactor(node)
	if math.Abs(float64(balanceFactor)) > 1 {
		return false
	}
	//if !atm.isBalanced(node.left) {
	//	return false
	//}else if !atm.isBalanced(node.right) {
	//	return false
	//}
	//return true
	return atm.isBalanced(node.left) && atm.isBalanced(node.right)
}
func (atm *AvlTreeMap) Add(k interface{}, v interface{}) {
	atm.root = atm.add(atm.root, k, v)
}
func (atm *AvlTreeMap) add(node *avlTreeMapNode, k interface{}, v interface{}) *avlTreeMapNode {
	if node == nil {
		atm.size++
		return &avlTreeMapNode{key: k.(Comparable), value: v, left: nil, right: nil, height: 1}
	}
	if k.(Comparable).Compare(node.key) < 0 {
		node.left = atm.add(node.left, k, v)
	} else if k.(Comparable).Compare(node.key) > 0 {
		node.right = atm.add(node.right, k, v)
	} else {
		node.value = v
		return node
	}

	node.height = 1 + int(math.Max(float64(atm.getHeight(node.left)), float64(atm.getHeight(node.right))))

	balance := atm.getBalanceFactor(node)

	if balance > 1 && atm.getBalanceFactor(node.left) >= 0 {
		//RR
		//fmt.Println(balance)
		return atm.rightRotate(node)
	}
	if balance < -1 && atm.getBalanceFactor(node.right) <= 0 {
		//LL
		return atm.leftRotate(node)
	}
	if balance > 1 && atm.getBalanceFactor(node.left) < 0 {
		//LR
		node.left = atm.leftRotate(node.left)
		return atm.rightRotate(node)
	}
	if balance < -1 && atm.getBalanceFactor(node.right) > 0 {
		//LL
		node.right = atm.rightRotate(node.right)
		return atm.leftRotate(node)
	}
	return node
}

/*
        y                  x          右旋转模式
       / \                / \
      x   t4            z    y
     / \               / \  / \
	z   t3           t1 t2 t3 t4
   / \
 t1  t2
*/
func (atm *AvlTreeMap) rightRotate(y *avlTreeMapNode) *avlTreeMapNode {
	x := y.left
	t3 := x.right

	x.right = y
	y.left = t3

	y.height = int(math.Max(float64(atm.getHeight(y.left)), float64(atm.getHeight(y.right)))) + 1
	x.height = int(math.Max(float64(atm.getHeight(x.left)), float64(atm.getHeight(x.right)))) + 1

	return x
}
func (atm *AvlTreeMap) leftRotate(y *avlTreeMapNode) *avlTreeMapNode {
	x := y.right
	t2 := x.left

	x.left = y
	y.right = t2

	y.height = int(math.Max(float64(atm.getHeight(y.left)), float64(atm.getHeight(y.right)))) + 1
	x.height = int(math.Max(float64(atm.getHeight(x.left)), float64(atm.getHeight(x.right)))) + 1

	return x
}
func (atm *AvlTreeMap) Remove(k interface{}) interface{} {
	ret := atm.Get(k)
	atm.root = atm.remove(atm.root, k)
	return ret
}
func (atm *AvlTreeMap) remove(node *avlTreeMapNode, k interface{}) *avlTreeMapNode {
	if node == nil {
		return nil
	}

	var retNode *avlTreeMapNode
	if k.(Comparable).Compare(node.key) < 0 {
		node.left = atm.remove(node.left, k)
		retNode = node
	} else if k.(Comparable).Compare(node.key) > 0 {
		node.right = atm.remove(node.right, k)
		retNode = node
	} else {
		if node.left == nil {
			right := node.right
			node.right = nil
			atm.size--
			retNode = right
		} else if node.right == nil {
			left := node.left
			node.left = nil
			atm.size--
			retNode = left
		} else {
			ret := atm.min(node.right)
			ret.right = atm.remove(node.right, ret.key)
			ret.left = node.left
			node.left, node.right = nil, nil
			retNode = ret
		}
	}
	if retNode == nil {
		return nil
	}
	retNode.height = 1 + int(math.Max(float64(atm.getHeight(retNode.left)), float64(atm.getHeight(retNode.right))))

	balance := atm.getBalanceFactor(retNode)

	if balance > 1 && atm.getBalanceFactor(retNode.left) >= 0 {
		//RR
		//fmt.Println(balance)
		return atm.rightRotate(retNode)
	}
	if balance < -1 && atm.getBalanceFactor(retNode.right) <= 0 {
		//LL
		return atm.leftRotate(retNode)
	}
	if balance > 1 && atm.getBalanceFactor(retNode.left) < 0 {
		//LR
		retNode.left = atm.leftRotate(retNode.left)
		return atm.rightRotate(retNode)
	}
	if balance < -1 && atm.getBalanceFactor(retNode.right) > 0 {
		//LL
		retNode.right = atm.rightRotate(retNode.right)
		return atm.leftRotate(retNode)
	}
	return retNode
}

func (atm *AvlTreeMap) min(node *avlTreeMapNode) *avlTreeMapNode {
	if node.left == nil {
		return node
	}
	return atm.min(node.left)
}
func (atm *AvlTreeMap) max(node *treeMapNode) *treeMapNode {
	if node.right == nil {
		return node
	}
	return atm.max(node.right)
}
func (atm *AvlTreeMap) Contains(key interface{}) bool {
	return atm.contains(atm.root, key)
}
func (atm *AvlTreeMap) contains(node *avlTreeMapNode, key interface{}) bool {
	if node == nil {
		return false
	} else if node.key == key {
		return true
	}
	if node.key.Compare(key.(Comparable)) < 0 {
		return atm.contains(node.left, key)
	} else {
		return atm.contains(node.right, key)
	}
}
func (atm *AvlTreeMap) Get(key interface{}) interface{} {
	return atm.get(atm.root, key)
}
func (atm *AvlTreeMap) get(node *avlTreeMapNode, key interface{}) interface{} {
	if atm.size == 0 {
		return nil
	}
	if node == nil {
		return nil
	} else if node.key == key {
		return node.value
	}

	if key.(Comparable).Compare(node.key) < 0 {
		return atm.get(node.left, key)
	} else {
		return atm.get(node.right, key)
	}
}
func (atm *AvlTreeMap) Set(key interface{}, newValue interface{}) {
	atm.set(atm.root, key, newValue)
}
func (atm *AvlTreeMap) set(node *avlTreeMapNode, key interface{}, newValue interface{}) {
	if atm.size == 0 {
		panic("bu cun zai zhe yang de key")
	}
	if node == nil {
		panic("bu cun zai zhe yang de key1")
	}
	if node.key == key {
		node.value = newValue
		return
	}

	if key.(Comparable).Compare(node.key) < 0 {
		atm.set(node.left, key, newValue)
	} else {
		atm.set(node.right, key, newValue)
	}
}
func (atm *AvlTreeMap) GetSize() int {
	return atm.size
}

func (atm *AvlTreeMap) IsEmpty() bool {
	return atm.size == 0
}

func (atm *AvlTreeMap) String() string {
	str := ""
	return atm.createString(atm.root, 0, str)
}

func (atm *AvlTreeMap) createString(node *avlTreeMapNode, depth int, str string) string {
	if node == nil {
		str += atm.depthString(depth) + "nil \n"
		return str
	}

	str += atm.depthString(depth) + "key:" + fmt.Sprint(node.key) + " value:" + fmt.Sprint(node.value) + "\n"
	str = atm.createString(node.left, depth+1, str)
	str = atm.createString(node.right, depth+1, str)
	return str
}
func (atm *AvlTreeMap) depthString(depth int) string {
	str := fmt.Sprint("")
	for i := 0; i < depth; i++ {
		str += fmt.Sprint("--")
	}
	return str
}
