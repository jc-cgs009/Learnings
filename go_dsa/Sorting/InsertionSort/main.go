package main

import "fmt"

func insertionSort(A []int) {
	n := len(A)
	for i := 0; i < n; i++ {
		cvalue := A[i]
		position := i
		for position > 0 && A[position-1] > cvalue {
			A[position] = A[position-1]
			position = position - 1
		}
		A[position] = cvalue
	}
}

func main() {
	A := []int{3, 5, 8, 9, 6, 2, -1, 0}
	fmt.Println("Original Array:", A)
	insertionSort(A)
	fmt.Println("Sorted Array:", A)
}
