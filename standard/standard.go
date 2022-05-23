package main

import (
	"bufio"
	. "fmt"
	"os"
)

func Solve() {
	//do something
}

func main() {
	T := 0 // 如果是单组样例，请讲 T := 0 改为 T := 1
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

// ri 将直接读入并返回一个 int
func ri() (x int) {
	scan(&x)
	return x
}

//rs 将直接读入并返回一个 string
func rs() (s string) {
	scan(&s)
	return s
}

//scan 用法与 fmt.scan() 相似，可以传入多个变量的地址进行读入操作
func scan(a ...interface{}) { Fscan(in, a...) }

//print 用法与 fmt.print() 相似，可以传入多个变量进行输出。注意，这个函数不会换行，多个变量之间有空格
func print(a ...interface{}) { Fprint(out, a...) }

//println 用法与 fmt.println() 相似，可以传入多个变量进行输出。这个函数会换行，多个变量之间有空格
func println(a ...interface{}) { Fprintln(out, a...) }

//min 可以传入多个 int 变量，求出这些变量的最小值
func min(a ...int) int {
	ans := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < ans {
			ans = a[i]
		}
	}
	return ans
}

//max 可以传入多个 int 变量，求出这些变量的最大值
func max(a ...int) int {
	ans := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > ans {
			ans = a[i]
		}
	}
	return ans
}

//sum 可以传入多个 int 变量，求出这些变量的总和，并返回总和。
//你可以使用 ... 运算符打散数组，使得 sum 函数可以求出数组的和。
func sum(a ...int) int {
	ans := 0
	for _, x := range a {
		ans += x
	}
	return ans
}

//pair 是一个数对，实际上就是 struct{x, y int}
type pair struct{ x, y int }
