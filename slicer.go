package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr []int // slicer

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		arr = append(arr, x) // dynamically append // --> likes vector
	}
	fmt.Println(arr)

}
