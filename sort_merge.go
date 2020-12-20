package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// マージソート：O(n log n) time complexity
func sort(a []int) []int {
	n := len(a)
	if n == 1 {
		return a
	}

	mid := n / 2

	// Step 1: 配列を二つに分ける
	var left []int
	var right []int
	for i := 0; i < n; i++ {
		if i < mid {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}

	// Step2: 再帰呼び出し（魔法）で、二つの配列をソートする
	left = sort(left)
	right = sort(right)

	// Step3: ソート済みの二つの配列をマージする
	// 番兵
	left = append(left, math.MaxInt32)
	right = append(right, math.MaxInt32)

	// 現在注目している、index
	lid := 0
	rid := 0

	b := make([]int, 0, n)

	for {
		if left[lid] == math.MaxInt32 && right[rid] == math.MaxInt32 {
			break
		}

		// 小さいほうを選んで、index を一つ進める
		if left[lid] < right[rid] {
			b = append(b, left[lid])
			lid++
		} else {
			b = append(b, right[rid])
			rid++
		}
	}

	return b
}

func main() {
	sc := NewScanner(os.Stdin)
	n := sc.Int()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = sc.Int()
	}

	a = sort(a)
	ans := 0
	for i := 0; i < n; i++ {
		ans += i*a[i] - (n-i-1)*a[i]
	}

	fmt.Println(ans)
}

// ----- 高速入力のためのテンプレ -----
type Scanner struct{ *bufio.Scanner }

func NewScanner(r io.Reader) *Scanner {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords) // 空白文字をデリミタとする(デフォルト)
	return &Scanner{Scanner: s}
}

// bufio.Scanner.Bytes はバッファの参照を返す(コピーではない)ため少し注意
func (s *Scanner) next() []byte {
	s.Scan()
	return s.Scanner.Bytes()
}

func (s *Scanner) Bytes() (b []byte) {
	unsafe := s.next()
	// コピーを取らないと上書きされる可能性がある
	b = make([]byte, len(unsafe))
	copy(b, unsafe)
	return
}

// バッファの参照に対しては、unsafe.Pointer は使えない
func (s *Scanner) Text() string { return string(s.next()) }

func (s *Scanner) Int() int {
	n, _ := strconv.Atoi(s.Text())
	return n
}
