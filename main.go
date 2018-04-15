package main

import (
	"fmt"
	"goSort/util"
	"os"
	"strings"
)

func main() {

	slice := util.GenSlice(10)
	algo := strings.ToLower(os.Args[1])

	util.Clear()

	fmt.Printf("Sorting algorithm: %v\n", algo)
	fmt.Printf("Data: %v Lenth: %v\n", slice, len(slice))

	// switch algo {
	// case "b", "bubble":
	// 	sort.Bubble(slice)
	// case "i", "insert", "insertion":
	// 	sort.Insertion(slice)
	// case "s", "select", "selection":
	// 	sort.Selection(slice)
	// case "m", "merge", "mergesort":
	// 	slice = sort.Merge(slice)
	// }

	fmt.Println("Sorted:", slice)
}
