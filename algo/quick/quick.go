package main

import "fmt"

func main() {
	arr := []int{8, 7, 9, 2, 3}
	QuickSort(arr)
	fmt.Println(arr)
}

func QuickSort(arr []int) {
	if len(arr) <= 1 { //base case
		return
	}

	//here i take the last number as pivot
	pivot := arr[len(arr)-1]
	j := 0
	i := 0
	for ; j < len(arr); j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	fmt.Println("out: ", arr)
	QuickSort(arr[:i])
	QuickSort(arr[i+1:])
}
