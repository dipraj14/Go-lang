package main

import (
	"fmt"
	"sort"
)

// sort the grap edges decresing order according ro their weight
type edges struct {
	x, y, w int
}

func main() {

	var e int
	fmt.Scan(&e)

	var adj []edges

	for i := 0; i < e; i++ {
		var x, y, w int
		fmt.Scan(&x, &y, &w)
		adj = append(adj, edges{x, y, w})
	}

	// 	for i := 0; i<e; i++{
	// 		fmt.Println(adj[i].w)
	// 	}

	sort.SliceStable(adj, func(i, j int) bool {
		return adj[i].w > adj[j].w
	})

	fmt.Println(adj)
}
