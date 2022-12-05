package main

import "fmt"

func binaryRecursiveSearch(A []int, key, s, e int) int {
	if s > e {
		return -1
	} else {
		mid := (s + e) / 2
		if A[mid] == key {
			return mid
		} else if key < A[mid] {
			return binaryRecursiveSearch(A, key, s, mid-1)
		} else if key > A[mid] {
			return binaryRecursiveSearch(A, key, mid+1, e)
		}
	}
	return -1
}

func main() {
	A := []int{10, 20, 56, 60, 89}
	found := binaryRecursiveSearch(A, 20, 0, len(A)-1)
	fmt.Println("Key found at inedex", found)
}
