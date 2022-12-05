package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (dl *DoublyLinkedList) len() int {
	return dl.size
}

func (dl *DoublyLinkedList) isempty() bool {
	return dl.size == 0
}

func (dl *DoublyLinkedList) addfirst(n int) {
	new := &Node{n, nil, nil}
	if dl.isempty() {
		// dl.head = new
		dl.tail = new
	} else {
		dl.head.prev = new
		new.next = dl.head
		// dl.head = new
	}
	dl.head = new
	dl.size++
}

func (dl *DoublyLinkedList) addlast(n int) {
	new := &Node{n, nil, nil}
	if dl.isempty() {
		dl.head = new
		// dl.tail = new
	} else {
		dl.tail.next = new
		new.prev = dl.tail
		// dl.tail = new
	}
	dl.tail = new
	dl.size++
}

func (dl *DoublyLinkedList) addany(n int, index int) {
	if index == 0 {
		dl.addfirst(n)
	} else if index >= (dl.len()) {
		dl.addlast(n)
	} else {
		new := &Node{n, nil, nil}
		p := dl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		new.prev = p
		new.next = p.next
		p.next.prev = new
		p.next = new
	}
	dl.size++
}

func (dl *DoublyLinkedList) removefirst() {
	if dl.isempty() {
		fmt.Println("Doubly Linked List is empty")
	} else {
		dl.head = dl.head.next
		dl.size--
	}
	if dl.isempty() {
		dl.head = nil
		dl.tail = nil
	}
}

func (dl *DoublyLinkedList) removelast() {
	if dl.isempty() {
		fmt.Println("Doubly linked list is empty")
	} else {
		dl.tail = dl.tail.prev
		dl.tail.next = nil
		dl.size--
	}
	if dl.isempty() {
		dl.head = nil
		dl.tail = nil
	}
}

func (dl *DoublyLinkedList) removeany(index int) {
	if index >= dl.len() {
		fmt.Println("Index out of range")
	} else if index == 0 {
		dl.removefirst()
	} else if index == (dl.len())-1 {
		dl.removelast()
	} else {
		p := dl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		p.next = p.next.next
		p.next.prev = p
		dl.size--
	}
	if dl.isempty() {
		dl.head = nil
		dl.tail = nil
	}
}

func (dl *DoublyLinkedList) display(n int) {
	if n == 0 {
		p := dl.head
		for p != nil {
			fmt.Print(p.data, "-->")
			p = p.next
		}
		fmt.Println()
	} else if n == 1 {
		p := dl.tail
		for p != nil {
			fmt.Print(p.data, "-->")
			p = p.prev
		}
		fmt.Println()
	} else {
		fmt.Println("Invalid argument: arugument must be 0 for forword display or 1 for reverse display")
	}
}

func main() {
	dll := DoublyLinkedList{}
	// 	dll.addfirst(10)
	// 	dll.addfirst(20)
	// 	dll.addfirst(30)
	// 	// dll.display(2)
	// 	dll.addlast(40)
	// 	dll.addlast(50)
	// 	dll.addlast(60)
	// 	dll.addany(11, 2)
	// 	// dll.display(0) // 0 --> forword diplay, 1--> reverse display
	// 	dll.removefirst()
	// 	dll.removefirst()
	// 	dll.display(0)
	// 	// dll.removelast()
	// 	// dll.removelast()
	// 	// dll.display(0)
	// 	// dll.removeany(dll.len() - 1)
	// 	dll.removeany(3)
	// 	dll.display(0)
	// 	fmt.Println(dll.len())
	dll.addfirst(10)
	dll.addfirst(20)
	dll.display(0)
	dll.removeany(0)
	dll.display(0)
}
