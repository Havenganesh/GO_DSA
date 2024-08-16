package main

import "fmt"

func main() {
	k := 6 //smallest
	arr := []int{8, 7, 9, 2, 3, 10}
	quickSelect(arr, k)
	// fmt.Println(arr)
}

func quickSelect(arr []int, key int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[len(arr)-1]
	i := 0

	for j := 0; j < len(arr); j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	if i < key-1 {
		key = key - i
		quickSelect(arr[i:], key)
	} else if i > key-1 {
		quickSelect(arr[:i], key)
	} else {
		fmt.Println(arr[i])
		return
	}

	// fmt.Println(arr)
}
