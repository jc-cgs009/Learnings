package main

import "fmt"

type sliceDequeue struct {
	dequeue []int
}

func (deq *sliceDequeue) len() int {
	return len(deq.dequeue)
}

func (deq *sliceDequeue) isempty() bool {
	return deq.len() == 0
}

func (deq *sliceDequeue) addfirst(n int) {
	if deq.isempty() {
		deq.dequeue = append(deq.dequeue, n)
	} else {
		t := []int{n}
		deq.dequeue = append(t, deq.dequeue...)
	}
}

func (deq *sliceDequeue) addlast(n int) {
	deq.dequeue = append(deq.dequeue, n)
}

func (deq *sliceDequeue) removefirst() int {
	r := deq.dequeue[0]
	deq.dequeue = deq.dequeue[1:]
	return r
}

func (deq *sliceDequeue) removelast() int {
	r := deq.dequeue[deq.len()-1]
	deq.dequeue = deq.dequeue[:deq.len()-1]
	return r
}

func (deq *sliceDequeue) first() int {
	return deq.dequeue[0]
}

// func (deq *sliceDequeue) last() int {
// 	return deq.dequeue[deq.last()-1]
// }

func (deq *sliceDequeue) display() {
	for _, v := range deq.dequeue {
		fmt.Print(v, "<-->")
	}
	fmt.Println()
}

func main() {
	dq := sliceDequeue{}
	fmt.Println(dq.len(), dq.isempty())
	dq.addfirst(10)
	dq.addfirst(20)
	dq.addfirst(30)
	dq.display()
	fmt.Println(dq.len(), dq.isempty())
	dq.addlast(40)
	dq.addlast(50)
	dq.addlast(60)
	dq.display()
	fmt.Println(dq.len(), dq.isempty())
	fmt.Println(dq.removefirst())
	dq.display()
	fmt.Println(dq.len(), dq.isempty())
	fmt.Println(dq.removelast())
	dq.display()
	fmt.Println(dq.len(), dq.isempty())
	fmt.Println(dq.first())
	// fmt.Println(dq.last())

}
