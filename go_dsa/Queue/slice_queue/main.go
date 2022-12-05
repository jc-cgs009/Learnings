package main

import "fmt"

type sliceQueue struct {
	queue []int
}

func (squ *sliceQueue) len() int {
	return len(squ.queue)
}

func (squ *sliceQueue) isempty() bool {
	return squ.len() == 0
}

func (squ *sliceQueue) enqueue(n int) {
	squ.queue = append(squ.queue, n)
}

func (squ *sliceQueue) dequeue() int {
	r := squ.queue[0]
	squ.queue = squ.queue[1:]
	return r
}

func (squ *sliceQueue) first() int {
	return squ.queue[0]
}

func (squ *sliceQueue) display() {
	for _, v := range squ.queue {
		fmt.Print(v, "<--")
	}
	fmt.Println()
}

func main() {
	sq := sliceQueue{}
	fmt.Println(sq.isempty())
	sq.enqueue(10)
	sq.enqueue(20)
	sq.enqueue(30)
	sq.display()
	fmt.Println(sq.len())
	fmt.Println(sq.isempty())
	fmt.Println(sq.dequeue())
	sq.display()
	fmt.Println(sq.len())
	fmt.Println(sq.isempty())
	fmt.Println(sq.first())
}
