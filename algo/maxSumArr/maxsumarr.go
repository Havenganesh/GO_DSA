package main

func main() {
	data := []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
	maxSumArr(data)
}

func maxSumArr(arr []int) {
	cum := 0
	maxSum := 0
	for i := 0; i < len(arr); i++ {
		println("cum ", cum)
		cum = cum + arr[i]
		if cum < 0 {
			cum = 0
		}
		if maxSum < cum {
			maxSum = cum
		}
	}
	println(maxSum)
}
