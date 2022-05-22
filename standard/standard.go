package standard

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

//===================================IO区===================================

//ri() 将直接读入并返回一个 int
func ri() (x int) {
	scan(&x)
	return x
}

//rs() 将直接读入并返回一个 string
func rs() (s string) {
	scan(&s)
	return s
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
