package main

import (
	"fmt"
)

func bucketSort(A []int) {
	n := len(A)
	maxelement := max(A)
	fmt.Println(maxelement)
	buckets := [10][]int{}
	for i := 0; i < n; i++ {
		j := int(n * A[i] / (maxelement + 1))
		buckets[j] = append(buckets[j], A[i])
	}

	for i := 0; i < 10; i++ {
		insertionSort(buckets[i])
	}

	k := 0

	for i := 0; i < 10; i++ {
		m := len(buckets[i])
		for j := 0; j < m; j++ {
			A[k] = buckets[i][j]
			k++
		}
	}

}

func max(A []int) int {
	i := 0
	j := len(A) - 1
	res := 0
	for i < j {
		if A[i] > A[j] {
			res = A[i]
			j--
		} else {
			res = A[j]
			i++
		}
	}
	return res
}

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
	bucketSort(A)
	fmt.Println("Sorted Array:", A)
}
