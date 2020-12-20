package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type UnionFind struct {
	n    int   // 頂点の数
	par  []int // 親の頂点番号。根の場合は -1
	size []int // 部分木のサイズ
}

func NewUnionFind(n int) *UnionFind {
	par := make([]int, n)
	size := make([]int, n)

	// 初期化
	// 始め、全ての頂点は根であり、サイズは 1 である
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
	// 部分木のサイズを返す
	return u.size[x]
}

func (u *UnionFind) Find(x int) int {

	// 根まで上に辿っていく
	for u.par[x] != -1 {
		x = u.par[x]
	}
	return x
}

func (u *UnionFind) Union(x, y int) {
	x = u.Find(x)
	y = u.Find(y)
	// すでに同じグループなら、何もしない
	if x == y {
		return
	}

	// サイズが小さい方から大きい方にマージする
	if x < y {
		x, y = y, x
	}
	// 根に辺を貼る
	u.par[y] = x
	// サイズを計算しなおす
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
