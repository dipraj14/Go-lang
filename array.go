package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n) // user define array
	arr1 := [4]int{3, 4, 5, 6}

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
	fmt.Println(arr1)

}
