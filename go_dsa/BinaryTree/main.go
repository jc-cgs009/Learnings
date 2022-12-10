package main

import (
	"fmt"
)

type Node struct {
	_element int
	_left    *Node
	_right   *Node
}

type binaryLinkedTree struct {
	_root *Node
}

func (bt *binaryLinkedTree) maketree(e int, left, right *binaryLinkedTree) {
	bt._root = &Node{e, left._root, right._root}
}

func (bt *binaryLinkedTree) inorder(troot *Node) {
	if troot != nil {
		bt.inorder(troot._left)
		fmt.Print(troot._element, " ")
		bt.inorder(troot._right)
	}
}

func (bt *binaryLinkedTree) preorder(troot *Node) {
	if troot != nil {
		fmt.Print(troot._element, " ")
		bt.preorder(troot._left)
		bt.preorder(troot._right)
	}
}

func (bt *binaryLinkedTree) postorder(troot *Node) {
	if troot != nil {
		bt.postorder(troot._left)
		bt.postorder(troot._right)
		fmt.Print(troot._element, " ")
	}
}

func (bt *binaryLinkedTree) levelorder(troot *Node) {
	Q := []*Node{}
	t := troot
	fmt.Print(troot._element, " ")
	Q = append(Q, t)
	for len(Q) > 0 {
		t = Q[0]
		Q = Q[1:]
		if t._left != nil {
			fmt.Print(t._left._element, " ")
			Q = append(Q, t._left)
		}
		if t._right != nil {
			fmt.Print(t._right._element, " ")
			Q = append(Q, t._right)
		}
	}
}

func (bt *binaryLinkedTree) count(troot *Node) int {
	if troot != nil {
		x := bt.count(troot._left)
		y := bt.count(troot._right)
		return x + y + 1
	}
	return 0
}

func (bt *binaryLinkedTree) height(troot *Node) int {
	if troot != nil {
		x := bt.height(troot._left)
		y := bt.height(troot._right)
		if x > y {
			return x + 1
		} else {
			return y + 1
		}
	}
	return 0
}

func main() {
	x := binaryLinkedTree{}
	y := binaryLinkedTree{}
	z := binaryLinkedTree{}
	r := binaryLinkedTree{}
	s := binaryLinkedTree{}
	t := binaryLinkedTree{}
	a := binaryLinkedTree{}
	x.maketree(40, &a, &a)
	y.maketree(60, &a, &a)
	z.maketree(20, &x, &a)
	r.maketree(50, &a, &y)
	s.maketree(30, &r, &a)
	t.maketree(10, &z, &s)
	fmt.Println("Inorder Traversal")
	t.inorder(t._root)
	fmt.Println()
	fmt.Println("Preorder Traversal")
	t.preorder(t._root)
	fmt.Println()
	fmt.Println("Postorder Traversal")
	t.postorder(t._root)
	fmt.Println()
	fmt.Println("Levelorder Traversal")
	t.levelorder(t._root)
	fmt.Println()
	fmt.Println("Number of Nodes :", t.count(t._root))
	fmt.Println("Height of the Tree :", t.height(t._root)-1)

}
