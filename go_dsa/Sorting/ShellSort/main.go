package main

import "fmt"

func shellSort(A []int) {
	n := len(A)
	gap := n / 2
	for gap > 0 {
		i := gap
		for i < n {
			gvalue := A[i]
			j := i - gap
			for j >= 0 && A[j] > gvalue {
				A[j+gap] = A[j]
				j = j - gap
			}
			A[j+gap] = gvalue
			i++
		}
		gap = gap / 2
	}
}

func main() {
	A := []int{3, 5, 8, 9, 6, 2}
	fmt.Println("Original Array:", A)
	shellSort(A)
	fmt.Println("Sorted Array:", A)
}
