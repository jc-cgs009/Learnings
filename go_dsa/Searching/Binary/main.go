package main

import "fmt"

func binarySearch(A []int, key int) int {
	s := 0
	e := len(A) - 1
	for s <= e {
		mid := (s + e) / 2
		if A[mid] == key {
			return mid
		} else if key < A[mid] {
			e = mid - 1
		} else if key > A[mid] {
			s = mid + 1
		}
	}
	return -1
}

func main() {
	A := []int{10, 20, 56, 60, 89}
	found := binarySearch(A, 60)
	fmt.Println("Key found at inedex", found)
}
