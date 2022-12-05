package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type stack struct {
	top  *Node
	size int
}

func (stk *stack) len() int {
	return stk.size
}

func (stk *stack) isempty() bool {
	return stk.size == 0
}

func (stk *stack) push(n int) {
	new := &Node{n, nil}
	if stk.isempty() {
		stk.top = new
	} else {
		new.next = stk.top
		stk.top = new
	}
	stk.size++
}

func (stk *stack) pop() int {
	var r int
	if stk.isempty() {
		fmt.Println("Stack is empty!")
	} else {
		r = stk.top.data
		stk.top = stk.top.next
		stk.size--
	}
	return r
}

func (stk *stack) Top() int {
	var r int
	if stk.isempty() {
		fmt.Println("Stack is empty!")
	} else {
		r = stk.top.data
	}
	return r
}

func (stk *stack) display() {
	p := stk.top
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
	fmt.Println()
}

func main() {
	s := stack{}
	s.push(10)
	s.push(20)
	s.push(30)
	s.display()
	fmt.Println(s.Top())
	fmt.Println(s.pop())
	s.display()
	// fmt.Println(s.Top())
}
