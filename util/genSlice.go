package util

import (
	"math/rand"
	"time"
)

// GenSlice returns slice of given size with random ints
func GenSlice(size int) []int {

	slice := make([]int, size, size)
	for i := range slice {

		rand.Seed(time.Now().UnixNano())
		slice[i] = rand.Intn(9) + 1
	}

	return slice
}
