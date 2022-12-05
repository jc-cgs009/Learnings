package main

import "fmt"

func quickSort(A []int, low, high int) {
	if low < high {
		pi := partition(A, low, high)
		quickSort(A, low, pi-1)
		quickSort(A, pi+1, high)
	}
}

func partition(A []int, low, high int) int {
	pivot := A[low]
	i := low + 1
	j := high
	for true {
		for i <= j && A[i] <= pivot {
			i++
		}

		for i <= j && A[j] > pivot {
			j--
		}

		if i <= j {
			A[i], A[j] = A[j], A[i]
		} else {
			break
		}
	}
	A[low], A[j] = A[j], A[low]
	return j

}

func main() {
	A := []int{3, 5, 6, 8, 9, 2, 0}
	fmt.Println("Original Array:", A)
	quickSort(A, 0, len(A)-1)
	fmt.Println("Sorted Array:", A)
}
