package main

import "fmt"

// pending
func main() {
	arr := []int{1, 2, 3, 4, 5}
	permutation(arr)
}

func permutation(arr []int) {
	for i := 1; i < len(arr); i++ {
		fmt.Println(arr[:i])
	}
}

func shuffler(arr []int) {

}
