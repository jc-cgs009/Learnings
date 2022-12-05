package main

import "fmt"

func linearSearch(A []int, key int) int {
	index := 0
	for index < len(A) {
		if A[index] == key {
			return index
		}
		index++
	}
	return -1
}

func main() {
	A := []int{12, 56, 84, 10, 5, 60}
	found := linearSearch(A, 12)
	fmt.Println("Key found at index", found)
}
