package main

import "fmt"

func main() {
	arr := []int{42, 17, 93, 58, 22, 6, 88, 15, 73, 4, 39, 50, 62, 91, 13, 37, 8, 72, 99, 31}
	insertionSort(arr)

}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ { //this is for loop through all the element
		current := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > current { // this is checking the element from j th position to 0 th index then go for the next i++ index
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}
	fmt.Println(arr)
}
