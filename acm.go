package main

import (
	"bufio"
	. "fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const MAXN = 2e5 + 5

func Solve() {
	var n int
	scan(&n)
	var a []int
	pre, suf := make([]int, n+1), make([]int, n+1)
	for i := 0; i < n; i++ {
		var x int
		scan(&x)
		a = append(a, x)
		suf[a[i]]++
	}
	var ans int64
	for i := 0; i < n; i++ {
		if i-1 >= 0 {
			pre[a[i-1]]++
		}
		suf[a[i]]--
		x := a[i]
		for j := max(x-2, 0); j <= min(x+2, n); j++ {
			for k := max(x-2, 0); k <= min(x+2, n); k++ {
				if max(x, max(j, k))-min(x, min(j, k)) <= 2 {
					ans += int64(pre[j]) * int64(suf[k])
				}
			}
		}
	}
	println(ans)
}

func main() {
	T := 0
	in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	if T == 0 {
		scan(&T)
	}
	for t := 0; t < T; t++ {
		Solve()
	}
	out.Flush()
}

var in *bufio.Reader
var out *bufio.Writer

func ri() (x int) {
	scan(&x)
	return x
}
func scan(a ...interface{})    { Fscan(in, a...) }
func print(a ...interface{})   { Fprint(out, a...) }
func println(a ...interface{}) { Fprintln(out, a...) }
