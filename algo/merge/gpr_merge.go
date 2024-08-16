package main

// import "fmt"

// func main() {
// 	arr := []int{8, 7, 9, 2, 3}
// 	mergeSort(arr)
// 	fmt.Println("final array : ", arr)
// }

// func mergeSort(arr []int) {
// 	if len(arr) <= 1 {
// 		return
// 	}
// 	middle := len(arr) / 2
// 	leftArr := make([]int, middle)
// 	rightArr := make([]int, len(arr)-middle)

// 	copy(leftArr, arr[:middle])
// 	copy(rightArr, arr[middle:])

// 	mergeSort(leftArr)
// 	mergeSort(rightArr)
// 	merge(leftArr, rightArr, arr)
// }

// func merge(leftArray []int, rightArray []int, array []int) {
// 	l := 0
// 	r := 0
// 	i := 0
// 	leftSize := len(leftArray)
// 	rightSize := len(rightArray)

// 	for l < leftSize && r < rightSize {
// 		if leftArray[l] < rightArray[r] {
// 			array[i] = leftArray[l]
// 			l++
// 		} else {
// 			array[i] = rightArray[r]
// 			r++
// 		}
// 		i++
// 	}

// 	for l < leftSize {
// 		array[i] = leftArray[l]
// 		i++
// 		l++
// 	}

// 	for r < rightSize {
// 		array[i] = rightArray[r]
// 		i++
// 		r++
// 	}
// }
