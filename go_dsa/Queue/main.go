package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type queue struct {
	front *Node
	rare  *Node
	size  int
}

func (qu *queue) len() int {
	return qu.size
}

func (qu *queue) isempty() bool {
	return qu.size == 0
}

func (qu *queue) enqueue(n int) {
	new := &Node{n, nil}
	if qu.isempty() {
		qu.front = new
	} else {
		qu.rare.next = new
	}
	qu.rare = new
	qu.size++
}

func (qu *queue) first() int {
	var r int
	if qu.isempty() {
		fmt.Println("Queue is empty!")
	} else {
		r = qu.front.data
	}
	return r
}

func (qu *queue) dequeue() int {
	var r int
	if qu.isempty() {
		fmt.Println("Queue is empty!")
	} else {
		r = qu.front.data
		qu.front = qu.front.next
		qu.size--
	}
	return r
}

func (qu *queue) display() {
	p := qu.front
	for p != nil {
		fmt.Print(p.data, "<--")
		p = p.next
	}
	fmt.Println()
}

func main() {
	q := queue{}
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.display()
	fmt.Println(q.len())
	fmt.Println(q.isempty())
	fmt.Println(q.dequeue())
	q.display()
	fmt.Println(q.len())
	fmt.Println(q.isempty())
	// q.dequeue()
	// q.dequeue()
	// q.dequeue()
	fmt.Println(q.first())

}
