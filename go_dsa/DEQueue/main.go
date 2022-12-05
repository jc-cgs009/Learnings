package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type DeQueue struct {
	front *Node
	rare  *Node
	size  int
}

func (dq *DeQueue) len() int {
	return dq.size
}

func (dq *DeQueue) isempty() bool {
	return dq.size == 0
}

func (dq *DeQueue) addfirst(n int) {
	new := &Node{n, nil}
	if dq.isempty() {
		dq.front = new
		dq.rare = new
	} else {
		new.next = dq.front
		dq.front = new
	}
	dq.size++
}

func (dq *DeQueue) addlast(n int) {
	new := &Node{n, nil}
	if dq.isempty() {
		dq.front = new
	} else {
		dq.rare.next = new
	}
	dq.rare = new
	dq.size++
}

func (dq *DeQueue) removefirst() int {
	var r int
	if dq.isempty() {
		fmt.Println("DEQueue is empty!")
	} else {
		r = dq.front.data
		dq.front = dq.front.next
		dq.size--
	}
	return r
}

func (dq *DeQueue) removelast() int {
	var r int
	if dq.isempty() {
		fmt.Println("DEQueue is empty!")
	} else {
		p := dq.front
		i := 1
		for i < dq.len()-1 {
			p = p.next
		}
		r = p.next.data
		p.next = nil
		dq.rare = p
		dq.size--
	}
	return r
}

func (dq *DeQueue) first() int {
	var r int
	if dq.isempty() {
		fmt.Println("DEQueue is empty!")
	} else {
		r = dq.front.data
	}
	return r
}

func (dq *DeQueue) last() int {
	var r int
	if dq.isempty() {
		fmt.Println("DEQueue is empty!")
	} else {
		r = dq.rare.data
	}
	return r
}

func (dq *DeQueue) display() {
	p := dq.front
	for p != nil {
		fmt.Print(p.data, "<-->")
		p = p.next
	}
	fmt.Println()
}

func main() {
	dqu := DeQueue{}
	dqu.addfirst(10)
	dqu.addfirst(20)
	dqu.display()
	fmt.Println(dqu.len())
	dqu.addlast(30)
	dqu.addlast(40)
	dqu.display()
	fmt.Println(dqu.len())
	fmt.Println(dqu.removefirst())
	dqu.display()
	fmt.Println(dqu.removefirst())
	dqu.display()
	fmt.Println(dqu.removelast())
	dqu.display()
	dqu.addfirst(10)
	dqu.addlast(20)
	dqu.display()
	fmt.Println(dqu.first())
	fmt.Println(dqu.last())
	dqu.addfirst(100)
	dqu.addlast(300)
	fmt.Println(dqu.first())
	fmt.Println(dqu.last())
}
