package main

// defer works on lifo rule --> last in fast out
import "fmt"

func main() {

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("zero") // fast print "zero" then 5 4 3 2 1 then "three" "two" "one"

	mydefer(6)
}
func mydefer(x int) {
	for i := 0; i < x; i++ {
		defer fmt.Println(i)
	}
}
