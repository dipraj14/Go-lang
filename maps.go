// count duplicate element using maps
// print the unique element
package main

import "fmt"

func count_dup(arr []int) map[int]int {
	dup_freq := make(map[int]int)

	// for index, val := range arr{ } // // index was not used so we use _ instade of index

	for _, value := range arr {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := dup_freq[value]

		if exist {
			dup_freq[value]++
		} else {
			dup_freq[value] = 1
		}
	}

	return dup_freq
}
func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	ans := count_dup(arr)

	fmt.Println(ans)

	for key, _ := range ans { // print all the unique element ... here value was not used so _ use instsde of value
		fmt.Println(key)
	}

}
