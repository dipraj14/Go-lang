package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	call_by_value(x)
	fmt.Println(x) // 4 does not changes in main funcation because of the copy pass
	call_by_reference(&x)
	fmt.Println(x) // 5 --> changes occour in main because of pass original value
}
func call_by_value(x int) {
	x++
}
func call_by_reference(x *int) {
	*x++ // derefence it
}
