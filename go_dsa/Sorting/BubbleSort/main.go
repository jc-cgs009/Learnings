package main

import "fmt"

func bubbleSort(A []int) {
	n := len(A)
	for passes := n - 1; passes > 0; passes-- {
		for i := 0; i < passes; i++ {
			if A[i] > A[i+1] {
				temp := A[i]
				A[i] = A[i+1]
				A[i+1] = temp
			}
		}
	}
}

func main() {
	A := []int{3, 5, 8, 9, 6, 2}
	fmt.Println("Original Array:", A)
	bubbleSort(A)
	fmt.Println("Sorted Array:", A)
}
