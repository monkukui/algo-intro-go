package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type UnionFind struct {
	n    int
	par  []int
	size []int
}

func NewUnionFind(n int) *UnionFind {
	par := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = -1
		size[i] = 1
	}
	return &UnionFind{
		n:    n,
		par:  par,
		size: size,
	}
}

func (u *UnionFind) Size(x int) int {
	return u.size[x]
}

func (u *UnionFind) Find(x int) int {
	for u.par[x] != -1 {
		x = u.par[x]
	}
	return x
}

func (u *UnionFind) Union(x, y int) {
	x = u.Find(x)
	y = u.Find(y)
	if x == y {
		return
	}
	if x < y {
		x, y = y, x
	}
	u.par[y] = x
	u.size[x] += u.size[y]
}

func main() {
	sc := NewScanner(os.Stdin)
	n, q := sc.Int(), sc.Int()

	uf := NewUnionFind(n)
	for i := 0; i < q; i++ {
		p, a, b := sc.Int(), sc.Int(), sc.Int()
		if p == 0 {
			uf.Union(a, b)
		} else {
			if uf.Find(a) == uf.Find(b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
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
