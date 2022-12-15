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

func (bt *binarySearchTree) insert(troot *Node, e int) {
	var temp *Node
	for troot != nil {
		temp = troot
		if e == troot.element {
			return
		} else if e < troot.element {
			troot = troot.left
		} else if e > troot.element {
			troot = troot.right
		}
	}

	n := &Node{e, nil, nil}
	if bt.root != nil {
		if e < temp.element {
			temp.left = n
		} else {
			temp.right = n
		}
	} else {
		bt.root = n
	}
}

func (bt *binarySearchTree) Search(key int) bool {
	troot := bt.root
	for troot != nil {
		if key == troot.element {
			return true
		} else if key < troot.element {
			troot = troot.left
		} else if key > troot.element {
			troot = troot.right
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
	B.insert(B.root, 50)
	B.insert(B.root, 30)
	B.insert(B.root, 80)
	B.insert(B.root, 10)
	B.insert(B.root, 40)
	B.insert(B.root, 60)
	B.insert(B.root, 90)
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
	fmt.Println(B.Search(50))
	fmt.Println(B.Search(100))
	B.delete(50)
	fmt.Println("Inorder Traversal")
	B.inorder(B.root)
	fmt.Println()

}
