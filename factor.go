package main

import "fmt"

func mul(a int, b int) int {
	return a * b
}

func main() {
	x := 5
	i := 1
	res := 1

	for i <= x {
		res = mul(res, i)
		i++
	}

	fmt.Println(res)
}
