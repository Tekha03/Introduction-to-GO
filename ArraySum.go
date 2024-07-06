package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5} 
	fmt.Println(ArraySum(arr, 5))
}

func ArraySum(arr []int, size int) int {
	sum := 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	return sum
}