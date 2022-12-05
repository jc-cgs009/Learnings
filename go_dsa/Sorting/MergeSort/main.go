package main

import "fmt"

func mergeSort(A []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(A, left, mid)
		mergeSort(A, mid+1, right)
		merge(A, left, mid, right)
	}
}

func merge(A []int, left, mid, right int) {
	i := left
	j := mid + 1
	k := left
	B := make([]int, (right + 1))

	for i <= mid && j <= right {
		if A[i] < A[j] {
			B[k] = A[i]
			i++
		} else {
			B[k] = A[j]
			j++
		}
		k++
	}

	for i <= mid {
		B[k] = A[i]
		i++
		k++
	}

	for j <= right {
		B[k] = A[j]
		j++
		k++
	}

	for x := left; x < right+1; x++ {
		A[x] = B[x]
	}
}

func main() {
	A := []int{3, 5, 8, 9, 6, 2}
	fmt.Println("Original Array:", A)
	mergeSort(A, 0, (len(A) - 1))
	fmt.Println("Sorted Array:", A)
}
