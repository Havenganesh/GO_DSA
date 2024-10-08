package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70}
	RotateArray(arr, 2)

}

func RotateArray(data []int, k int) {
	n := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)
}

func ReverseArray(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	fmt.Println(data)
}
