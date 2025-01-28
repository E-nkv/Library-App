package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func genTagsRandomly(et []int64) []int64 {
	var idxs []int
	var res []int64

	N := len(et)
	amount := rand.IntN(11)
	for range amount {
		var n int
		n = rand.IntN(N)

		for slices.Contains(idxs, n) {

			n = rand.IntN(N)
		}
		idxs = append(idxs, n)
		res = append(res, et[n])

	}

	return res
}

func main() {
	et := []int64{56, 1, 95, 6466, 231, 75, 42, 74, 43, 66, 99}
	m := map[int]int{}
	for i := 0; i < 10000; i++ {
		res := genTagsRandomly(et)
		m[len(res)]++
	}
	fmt.Println(m)
}
