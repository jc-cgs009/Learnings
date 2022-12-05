package main

import "fmt"

func selectionSort(A []int) {
	n := len(A)
	for i := 0; i < n; i++ {
		position := i
		for j := i; j < n; j++ {
			if A[position] > A[j] {
				position = j
			}
		}
		temp := A[i]
		A[i] = A[position]
		A[position] = temp
	}
}

func main() {
	A := []int{3, 5, 8, 9, 6, 2}
	fmt.Println("Original Array:", A)
	selectionSort(A)
	fmt.Println("Sorted Array:", A)
}
