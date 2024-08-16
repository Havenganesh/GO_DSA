package main

import "fmt"

func main() {
	arr := []int{8, 7, 9, 2, 3}
	selectionSort(arr)

}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	fmt.Println(arr)
}
