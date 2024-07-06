package main

import "fmt"

func main() {
	var length int
	fmt.Scan(&length)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		var x int
		fmt.Scan(&x)
		arr[i] = x
	}
	var n int
	fmt.Scan(&n)
	if Contains(arr, length, n) {
		fmt.Println("Contains")
	} else {
		fmt.Println("Doesn't contain")
	}
}

func Contains(arr []int, length int, n int) bool {
	for i := 0; i < length; i++ {
		if arr[i] == n {
			return true
		}
	}
	return false
}