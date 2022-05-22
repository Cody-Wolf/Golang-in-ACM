package main

import (
	"bufio"
	. "fmt"
	"os"
)

func Solve() {
	n := ri()
	var s string
	scan(&s)
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			cnt++
		}
	}
	if cnt%2 == 1 {
		println("NO")
		return
	}
	println("YES")
	type pair struct{ l, r int }
	var p []pair
	for i := 0; i < n; i++ {
		j := i + 1
		for s[j] == '0' {
			j++
		}
		p = append(p, pair{i, j})
		i = j
	}
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, i)
	}
	for i := 0; i < len(p)-1; i++ {
		for j := p[i].l; j < p[i].r; j++ {
			println(j+1, j+2)
			a[j], a[j+1] = p[i].l, p[i].l
		}
	}
	var b []int
	b = append(b, a[0])
	for i := 1; i < n; i++ {
		if a[i] != a[i-1] {
			b = append(b, a[i])
		}
	}
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

//===================================IO区===================================

//ri() 将直接读入并返回一个 int
func ri() (x int) {
	scan(&x)
	return x
}

//scan() 用法与 fmt.scan() 相似，可以传入多个变量的地址进行读入操作
func scan(a ...interface{}) { Fscan(in, a...) }

func print(a ...interface{}) { Fprint(out, a...) }

func println(a ...interface{}) { Fprintln(out, a...) }

//===================================IO区===================================

//===================================基础区===================================

func min(a ...int) int {
	ans := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < ans {
			ans = a[i]
		}
	}
	return ans
}

func max(a ...int) int {
	ans := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > ans {
			ans = a[i]
		}
	}
	return ans
}

//===================================基础区===================================
