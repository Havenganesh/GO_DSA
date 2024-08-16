package main

import "fmt"

func main() {
	// Initialize an array with 20 hardcoded values
	arr := []int{42, 17, 93, 58, 22, 6, 88, 15, 73, 4, 39, 50, 62, 91, 13, 37, 8, 72, 99, 31}
	// arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	BubbleSort(arr)
}

func BubbleSort(arr []int) {
	fmt.Println(arr)
	c := 0
	for i := 0; i < len(arr); i++ {
		c++
		swap := false
		for j := 0; j < len(arr)-i-1; j++ {
			c++
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	fmt.Println("final", arr)
	fmt.Println("count : ", c, " size ", len(arr))
}
