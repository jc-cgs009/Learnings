package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyCircular struct {
	head *Node
	tail *Node
	size int
}

func (dc *DoublyCircular) len() int {
	return dc.size
}

func (dc *DoublyCircular) isempty() bool {
	return dc.size == 0
}

func (dc *DoublyCircular) addfirst(n int) {
	new := &Node{n, nil, nil}
	if dc.isempty() {
		dc.head = new
		dc.tail = new
		dc.tail.next = dc.head
		dc.head.prev = dc.tail
	} else {
		new.next = dc.head
		new.prev = dc.head.prev
		dc.head.prev = new
		dc.head = new
		dc.tail.next = dc.head
	}
	dc.size++
}

func (dc *DoublyCircular) addlast(n int) {
	new := &Node{n, nil, nil}
	if dc.isempty() {
		dc.head = new
		dc.tail = new
		dc.tail.next = dc.head
		dc.head.prev = dc.tail
	} else {
		new.prev = dc.tail
		new.next = dc.tail.next
		dc.tail.next = new
		dc.tail = new
		dc.head.prev = dc.tail
	}
	dc.size++
}

func (dc *DoublyCircular) addany(n, i int) {
	if i == 0 {
		dc.addfirst(n)
	} else if i >= dc.len() {
		dc.addlast(n)
	} else {
		new := &Node{n, nil, nil}
		p := dc.head
		x := 1
		for x < i {
			p = p.next
			x++
		}
		new.prev = p
		new.next = p.next
		p.next = new
		dc.size++
	}
}

func (dc *DoublyCircular) removefirst() {
	if dc.isempty() {
		fmt.Println("Doubly circular linked list is empty!")
	} else {
		dc.head = dc.head.next
		dc.head.prev = dc.tail
		dc.tail.next = dc.head
		dc.size--
	}
	if dc.isempty() {
		dc.head = nil
		dc.tail = nil
	}
}

func (dc *DoublyCircular) removelast() {
	if dc.isempty() {
		fmt.Println("Douly circular linked list is empty!")
	} else {
		dc.tail = dc.tail.prev
		dc.tail.next = dc.head
		dc.head.prev = dc.tail
		dc.size--
	}
	if dc.isempty() {
		dc.head = nil
		dc.tail = nil
	}
}

func (dc *DoublyCircular) removeany(i int) {
	if i == 0 {
		dc.removefirst()
	} else if i == dc.len()-1 {
		dc.removelast()
	} else if i >= dc.len() {
		fmt.Println("List index out of range!")
	} else {
		p := dc.head
		x := 1
		for x < i {
			p = p.next
			x++
		}
		p.next = p.next.next
		p.next.prev = p
		dc.size--
	}
}

func (dc *DoublyCircular) display() {
	p := dc.head
	i := 0
	for i < dc.size {
		fmt.Print(p.data, "-->")
		p = p.next
		i++
	}
	fmt.Println()
}

func main() {
	dcl := DoublyCircular{}
	dcl.addfirst(10)
	dcl.addfirst(20)
	dcl.addfirst(30)
	dcl.display()
	fmt.Println(dcl.size)
	dcl.addlast(10)
	dcl.addlast(20)
	dcl.addlast(30)
	dcl.display()
	fmt.Println(dcl.size)
	fmt.Println(dcl.head, dcl.tail.next)
	dcl.addany(100, 6)
	dcl.display()
	fmt.Println(dcl.size)
	fmt.Println(dcl.head, dcl.tail.next)
	fmt.Println(dcl.head.prev, dcl.tail)
	dcl.removefirst()
	dcl.display()
	fmt.Println(dcl.head, dcl.tail.next)
	fmt.Println(dcl.head.prev, dcl.tail)
	dcl.removelast()
	dcl.display()
	fmt.Println(dcl.head, dcl.tail.next)
	fmt.Println(dcl.head.prev, dcl.tail)
	dcl.removeany(3)
	dcl.display()
	fmt.Println(dcl.head, dcl.tail.next)
	fmt.Println(dcl.head.prev, dcl.tail)

}
