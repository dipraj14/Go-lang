// Concurrency and goroutines

package main

import (
	"fmt"
	"time"
)

func main() {
	go routin("hello") // this task was skiped or forget
	routin("world")
}
func routin(s string) {
	i := 0
	for ; i < 6; i++ {
		time.Sleep(3 * time.Microsecond) // when Sleep the forget task was excuted
		fmt.Println(s)
	}
}

//https://www.youtube.com/watch?v=V-0ifUKCkBI&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=52
