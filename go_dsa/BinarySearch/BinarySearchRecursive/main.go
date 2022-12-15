package main

import "fmt"

type Node struct {
	element int
	left    *Node
	right   *Node
}

type binarySearchTree struct {
	root *Node
}

func (bt *binarySearchTree) rinsert(troot *Node, e int) *Node {
	if troot != nil {
		if e < troot.element {
			troot.left = bt.rinsert(troot.left, e)
		} else if e > troot.element {
			troot.right = bt.rinsert(troot.right, e)
		}
	} else {
		n := &Node{e, nil, nil}
		troot = n
	}
	return troot
}

func (bt *binarySearchTree) rSearch(troot *Node, key int) bool {
	if troot != nil {
		if key == troot.element {
			return true
		} else if key < troot.element {
			return bt.rSearch(troot.left, key)
		} else if key > troot.element {
			return bt.rSearch(troot.right, key)
		}
	}
	return false
}

func (bt *binarySearchTree) delete(e int) {
	var pp *Node
	p := bt.root
	for p != nil && p.element != e {
		pp = p
		if e < p.element {
			p = p.left
		} else {
			p = p.right
		}
	}

	if p == nil {
		return
	}

	if p.left != nil && p.right != nil {
		s := p.left
		ps := p
		for s.right != nil {
			ps = s
			s = s.right
		}
		p.element = s.element
		p = s
		pp = ps
	}

	var c *Node
	if p.left != nil {
		c = p.left
	} else {
		c = p.right
	}

	if p == bt.root {
		bt.root = nil
	} else {
		if p == pp.left {
			pp.left = c
		} else {
			pp.right = c
		}
	}

}

func (bt *binarySearchTree) inorder(troot *Node) {
	if troot != nil {
		bt.inorder(troot.left)
		fmt.Print(troot.element, " ")
		bt.inorder(troot.right)
	}
}

func (bt *binarySearchTree) preorder(troot *Node) {
	if troot != nil {
		fmt.Print(troot.element, " ")
		bt.preorder(troot.left)
		bt.preorder(troot.right)
	}
}

func (bt *binarySearchTree) postorder(troot *Node) {
	if troot != nil {
		bt.postorder(troot.left)
		bt.postorder(troot.right)
		fmt.Print(troot.element, " ")
	}
}

func (bt *binarySearchTree) levelorder(troot *Node) {
	Q := []*Node{}
	t := troot
	fmt.Print(t.element, " ")
	Q = append(Q, t)
	for len(Q) > 0 {
		t = Q[0]
		Q = Q[1:]
		if t.left != nil {
			fmt.Print(t.left.element, " ")
			Q = append(Q, t.left)
		}
		if t.right != nil {
			fmt.Print(t.right.element, " ")
			Q = append(Q, t.right)
		}
	}
}

func main() {
	B := binarySearchTree{}
	B.root = B.rinsert(B.root, 50)
	B.rinsert(B.root, 30)
	B.rinsert(B.root, 80)
	B.rinsert(B.root, 10)
	B.rinsert(B.root, 40)
	B.rinsert(B.root, 60)
	B.rinsert(B.root, 90)
	fmt.Println("Inorder Traversal")
	B.inorder(B.root)
	fmt.Println()
	fmt.Println("Preorder Traversal")
	B.preorder(B.root)
	fmt.Println()
	fmt.Println("Postorder Traversal")
	B.postorder(B.root)
	fmt.Println()
	fmt.Println("Levelorder Traversal")
	B.levelorder(B.root)
	fmt.Println()
	fmt.Println(B.rSearch(B.root, 50))
	fmt.Println(B.rSearch(B.root, 100))
	B.delete(50)
	fmt.Println("Inorder Traversal")
	B.inorder(B.root)
	fmt.Println()

}
