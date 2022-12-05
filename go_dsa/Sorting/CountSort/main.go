package main

import "fmt"

func countSort(A []int) {
	n := len(A)
	maxnum := A[0]
	for _, v := range A {
		if v > maxnum {
			maxnum = v
		}
	}

	carray := make([]int, maxnum+1)

	for i := 0; i < n; i++ {
		carray[A[i]] = carray[A[i]] + 1
	}

	i, j := 0, 0
	for j < maxnum+1 {
		if carray[j] > 0 {
			A[i] = j
			carray[j] -= 1
			i++
		} else {
			j++
		}
	}
}

func main() {
	A := []int{3, 5, 8, 9, 6, 2, 3, 5, 5}
	fmt.Println("Original Array:", A)
	countSort(A)
	fmt.Println("Sorted Array:", A)
}
