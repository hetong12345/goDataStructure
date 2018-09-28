package data

import (
	"fmt"
)

type BinarySearchTree struct {
	root *bsTreeNode
	size int
}

type bsTreeNode struct {
	value       Comparable
	left, right *bsTreeNode
}

func CreateBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
		size: 0,
	}
}

func (bst *BinarySearchTree) GetSize() int {
	return bst.size
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.size == 0
}

func (bst *BinarySearchTree) Add(value Comparable) {
	bst.root = bst.add(bst.root, value)
}
func (bst *BinarySearchTree) add(node *bsTreeNode, value Comparable) *bsTreeNode {
	if node == nil {
		bst.size++
		return &bsTreeNode{value, nil, nil}
	}

	if value.Compare(node.value) < 0 {
		node.left = bst.add(node.left, value)
	} else if value.Compare(node.value) > 0 {
		node.right = bst.add(node.right, value)
	}
	return node
}
func (bst *BinarySearchTree) Contains(value Comparable) bool {
	return bst.contains(bst.root, value)
}
func (bst *BinarySearchTree) contains(node *bsTreeNode, value Comparable) bool {
	if node == nil {
		return false
	} else if node.value == value {
		return true
	}
	if value.Compare(node.value) < 0 {
		return bst.contains(node.left, value)
	} else {
		return bst.contains(node.right, value)
	}
}

func (bst *BinarySearchTree) PreOrder() { //前序遍历
	bst.preOrder(bst.root)
}
func (bst *BinarySearchTree) preOrder(node *bsTreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	bst.preOrder(node.left)
	bst.preOrder(node.right)
}

func (bst *BinarySearchTree) MidOrder() { //中序遍历
	bst.midOrder(bst.root)
}
func (bst *BinarySearchTree) midOrder(node *bsTreeNode) {
	if node == nil {
		return
	}
	bst.midOrder(node.left)
	fmt.Println(node.value)
	bst.midOrder(node.right)
}
func (bst *BinarySearchTree) PostOrder() { //后序遍历
	bst.postOrder(bst.root)
}
func (bst *BinarySearchTree) postOrder(node *bsTreeNode) {
	if node == nil {
		return
	}
	bst.postOrder(node.left)
	bst.postOrder(node.right)
	fmt.Println(node.value)
}
func (bst *BinarySearchTree) LevelOrder() {
	lq := CreateLoopQueue(10)
	lq.EnQueue(bst.root)
	for !lq.IsEmpty() {
		cur := lq.DeQueue().(*bsTreeNode)
		fmt.Println(cur.value)
		if cur.left != nil {
			lq.EnQueue(cur.left)
		}
		if cur.right != nil {
			lq.EnQueue(cur.right)
		}
	}
}
func (bst *BinarySearchTree) Min() Comparable { //找出最小值
	if bst.size == 0 {
		return nil
	}
	return bst.min(bst.root).value
}
func (bst *BinarySearchTree) min(node *bsTreeNode) *bsTreeNode {
	if node.left == nil {
		return node
	}
	return bst.min(node.left)
}

func (bst *BinarySearchTree) RemoveMin() Comparable {
	ret := bst.min(bst.root)
	bst.root = bst.removeMin(bst.root)
	bst.size--
	return ret.value
}
func (bst *BinarySearchTree) removeMin(node *bsTreeNode) *bsTreeNode {
	if node.left == nil {
		right := node.right
		node.right = nil
		return right
	}

	node.left = bst.removeMin(node.left)
	return node
}
func (bst *BinarySearchTree) Max() Comparable { //找出最大值
	if bst.size == 0 {
		return nil
	}
	return bst.max(bst.root).value
}
func (bst *BinarySearchTree) max(node *bsTreeNode) *bsTreeNode {
	if node.right == nil {
		return node
	}
	return bst.max(node.right)
}
func (bst *BinarySearchTree) RemoveMax() Comparable {
	ret := bst.max(bst.root)
	bst.root = bst.removeMax(bst.root)
	bst.size--
	return ret.value
}
func (bst *BinarySearchTree) removeMax(node *bsTreeNode) *bsTreeNode {
	if node.right == nil {
		left := node.left
		node.right = nil
		return left
	}

	node.right = bst.removeMax(node.right)
	return node
}

func (bst *BinarySearchTree) Remove(value Comparable) {
	bst.root = bst.remove(bst.root, value)
}
func (bst *BinarySearchTree) remove(node *bsTreeNode, value Comparable) *bsTreeNode {
	if node == nil {
		return nil
	}
	if value.Compare(node.value) < 0 {
		node.left = bst.remove(node.left, value)
		return node
	} else if value.Compare(node.value) > 0 {
		node.right = bst.remove(node.right, value)
		return node
	} else {
		if node.left == nil {
			right := node.right
			node.right = nil
			bst.size--
			return right
		} else if node.right == nil {
			left := node.left
			node.left = nil
			bst.size--
			return left
		}

		ret := bst.min(node.right)
		ret.right = bst.removeMin(node.right)
		ret.left = node.left
		node.left, node.right = nil, nil
		return ret

	}

}
func (bst *BinarySearchTree) String() string {
	str := ""
	return bst.createString(bst.root, 0, str)
}
func (bst *BinarySearchTree) createString(node *bsTreeNode, depth int, str string) string {
	if node == nil {
		str += bst.depthString(depth) + "nil \n"
		return str
	}

	str += bst.depthString(depth) + fmt.Sprint(node.value) + "\n"
	str = bst.createString(node.left, depth+1, str)
	str = bst.createString(node.right, depth+1, str)
	return str
}
func (bst *BinarySearchTree) depthString(depth int) string {
	str := fmt.Sprint("")
	for i := 0; i < depth; i++ {
		str += fmt.Sprint("--")
	}
	return str
}
