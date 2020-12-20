package main

import (
	"fmt"
	"math"
)

// ナイーブソート： O(n) time complexity
func sort(a []int) []int {
	n := len(a)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		// min: 最小値, id: 最小値の index
		min := math.MaxInt32
		id := -1

		for j := 0; j < n; j++ {
			if min > a[j] {
				min = a[j]
				id = j
			}
		}

		// 選択済みにする
		a[id] = math.MaxInt32

		// ソート済み配列に追加
		b[i] = min
	}

	return b
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	a = sort(a)
	fmt.Println(a)
}
