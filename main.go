package main

import "fmt"

func main() {
	// Arrays have fixed size
	arr := [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Println("Array:", arr) // Output: Array: [1 2 3] - original array is unchanged

	// Slices are dynamic
	sl := []int{1, 2, 3}
	modifySlice(sl)
	fmt.Println("Slice:", sl) // Output: Slice: [100 2 3] - original slice is modified
}

func modifyArray(arr [3]int) {
	arr[0] = 100
}

func modifySlice(sl []int) {
	sl[0] = 100
}
