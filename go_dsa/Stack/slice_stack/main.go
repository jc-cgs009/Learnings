package main

import "fmt"

type stackSlice struct {
	stack []int
}

func (sst *stackSlice) len() int {
	return len(sst.stack)
}

func (sst *stackSlice) isempty() bool {
	return len(sst.stack) == 0
}

func (sst *stackSlice) push(n int) {
	sst.stack = append(sst.stack, n)
}

func (sst *stackSlice) pop() int {
	r := sst.stack[sst.len()-1]
	sst.stack = sst.stack[:sst.len()-1]
	return r
}

func (sst *stackSlice) Top() int {
	return sst.stack[sst.len()-1]
}

func (sst *stackSlice) display() {
	for _, v := range sst.stack {
		fmt.Print(v, "<--")
	}
	fmt.Println()
}

func main() {
	ss := stackSlice{}
	ss.push(10)
	ss.push(20)
	ss.push(30)
	ss.display()
	fmt.Println(ss.pop())
	ss.display()
	// fmt.Println(ss.pop())
	// ss.display()
	fmt.Println(ss.Top())
}
