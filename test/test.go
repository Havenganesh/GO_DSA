// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	arr := []int{8, 7, 9, 2, 3}
	mergeSort(arr)
}

func mergeSort(arr []int) {
	size := len(arr)
	if size <= 1 {
		return
	}
	middle := size / 2
	leftArr := make([]int, middle)
	rightArr := make([]int, size-middle)
	copy(leftArr, arr[:middle])
	copy(rightArr, arr[middle:])
	//fmt.Println(leftArr, rightArr)
	mergeSort(leftArr)
	mergeSort(rightArr)
	merge(leftArr, rightArr, arr)
	fmt.Println(arr)
}

func merge(leftArr []int, rightArr []int, array []int) {
	l, r, i := 0, 0, 0
	lsize := len(leftArr)
	rsize := len(rightArr)
	for l < lsize && r < rsize {

		if leftArr[l] < rightArr[r] {
			array[i] = leftArr[l]
			i++
			l++
		} else {
			array[i] = rightArr[r]
			i++
			r++
		}
	}

	for l < lsize {
		array[i] = leftArr[l]
		l++
		i++
	}

	for r < rsize {
		array[i] = rightArr[r]
		r++
		i++
	}
}
