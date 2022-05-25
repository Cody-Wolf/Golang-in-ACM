//template from https://github.com/Cody-Wolf/Golang-in-ACM
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
	T := 1 // 如果是多组样例，请讲 T := 1 改为 T := 0
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

// ri64 将直接读入并返回一个 int64
func ri64() (x int64) {
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

//print 用法与 fmt.print() 相似，可以传入多个变量进行输出。
func print(a ...interface{}) { Fprint(out, a...) }

//println 用法与 fmt.println() 相似，可以传入多个变量进行输出并在最后换行。
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

//min64 可以传入多个 int64 变量，求出这些变量的最小值
func min64(a ...int64) int64 {
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

//max64 可以传入多个 int64 变量，求出这些变量的最大值
func max64(a ...int64) int64 {
	ans := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > ans {
			ans = a[i]
		}
	}
	return ans
}

//sum 可以传入多个 int 变量，求出这些变量的总和，并返回总和。
//
//你可以使用 ... 运算符打散数组，使得 sum 函数可以求出数组的和。
func sum(a ...int) int {
	ans := 0
	for _, x := range a {
		ans += x
	}
	return ans
}

//sum64 可以传入多个 int64 变量，求出这些变量的总和，并返回总和。
//
//你可以使用 ... 运算符打散数组，使得 sum64 函数可以求出数组的和。
func sum64(a ...int64) int64 {
	ans := int64(0)
	for _, x := range a {
		ans += x
	}
	return ans
}

//reverse 用于翻转 int 切片，所有元素将会变成原切片的倒序
func reverse(a []int) {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
}

//reverse64 用于翻转 int64 切片，所有元素将会变成原切片的倒序
func reverse64(a []int64) {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
}

// lowerBound 类似于 C++ 中的 lower_bound，用于二分查找 int 切片中，第一个大于等于 k 的元素在切片中的位置。
//
//如果元素不存在，则返回切片的长度。
//
//注意，lowerBound 作用的 int 切片必须升序。
func lowerBound(a []int, k int) int {
	l, r, pos := 0, len(a)-1, len(a)
	for l <= r {
		mid := (l + r) >> 1
		if a[mid] < k {
			l = mid + 1
		} else {
			pos, r = mid, mid-1
		}
	}
	return pos
}

//lowerBound64 类似于 C++ 中的 lower_bound，用于二分查找 int64 切片中，第一个大于等于 k 的元素在切片中的位置。
//
//如果元素不存在，则返回切片的长度。
//
//注意，lowerBound64 作用的 int64 切片必须升序。
func lowerBound64(a []int64, k int64) int {
	l, r, pos := 0, len(a)-1, len(a)
	for l <= r {
		mid := (l + r) >> 1
		if a[mid] < k {
			l = mid + 1
		} else {
			pos, r = mid, mid-1
		}
	}
	return pos
}

//upperBound 类似于 C++ 中的 upper_bound，用于二分查找 int 切片中，第一个大于 k 的元素在切片中的位置。
//
//如果元素不存在，则返回切片的长度。
//
//注意，upperBound 作用的 int 切片必须升序。
func upperBound(a []int, k int) int {
	l, r, pos := 0, len(a)-1, len(a)
	for l <= r {
		mid := (l + r) >> 1
		if a[mid] <= k {
			l = mid + 1
		} else {
			pos, r = mid, mid-1
		}
	}
	return pos
}

//upperBound64 类似于 C++ 中的 upper_bound，用于二分查找 int64 切片中，第一个大于 k 的元素在切片中的位置。
//
//如果元素不存在，则返回切片的长度。
//
//注意，upperBound64 作用的 int64 切片必须升序。
func upperBound64(a []int64, k int64) int {
	l, r, pos := 0, len(a)-1, len(a)
	for l <= r {
		mid := (l + r) >> 1
		if a[mid] <= k {
			l = mid + 1
		} else {
			pos, r = mid, mid-1
		}
	}
	return pos
}

//unique 类似于 C++ 中的 unique，用于给切片去重。
//
//区别在于: unique 函数接受一个 int 切片，并返回一个新的、去好重的切片
//
//注意: 传入的切片必须升序
func unique(a []int) []int {
	var b []int
	b = append(b, a[0])
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			b = append(b, a[i])
		}
	}
	return b
}

//valSet 将传入切片的所有元素，统一赋值为传入的元素 k
func valSet(a []int, k int) {
	for i, _ := range a {
		a[i] = k
	}
}

//Stack 类型是一个 int 类型的栈
type Stack struct {
	s []int
}

//top 获取栈顶元素
func (st *Stack) top() int { return st.s[len(st.s)-1] }

//size 返回栈内的元素个数
func (st *Stack) size() int { return len(st.s) }

//empty 返回一个 bool 类型，表示栈是否为空
func (st *Stack) empty() bool { return st.size() == 0 }

//pop 弹出栈顶元素
func (st *Stack) pop() { st.s = st.s[:len(st.s)-1] }

//push 将一个 int 类型从栈顶添加入栈
func (st *Stack) push(x int) { st.s = append(st.s, x) }

//get 得到栈的内容（一个新的 int切片）
func (st *Stack) get() (ans []int) {
	copy(ans, st.s)
	return ans
}

//Queue 类型是一个 int 类型的队列
type Queue struct {
	q    []int
	head int
}

//front 获取队首元素
func (que *Queue) front() int { return que.q[que.head] }

//size 返回队列的元素个数
func (que *Queue) size() int { return len(que.q) - que.head }

//empty 返回一个 bool 类型，表示队列是否为空
func (que *Queue) empty() bool { return que.size() == 0 }

//pop 弹出队首元素
func (que *Queue) pop() { que.head++ }

//push 将一个 int 类型从队尾添加入队
func (que *Queue) push(x int) { que.q = append(que.q, x) }

//get 得到队列的内容（一个新的 int切片）
func (que *Queue) get() (ans []int) {
	copy(ans, que.q)
	return ans
}

//PriorityQueue 类似是一个 int 型的优先队列
//
//注意：你需要通过 newPriorityQueue 函数来获取一个优先队列实例，而不是直接创建
type PriorityQueue struct {
	heap []int
	cmp  func(int, int) bool
}

//heapifyUp 从 x 结点开始到堆顶，调整结点使得满足堆的性质
//
//注意：你不应该直接使用这个函数
func (pq *PriorityQueue) heapifyUp(x int) {
	for x > 1 && pq.cmp(pq.heap[x-1], pq.heap[x/2-1]) {
		pq.heap[x-1], pq.heap[x/2-1] = pq.heap[x/2-1], pq.heap[x-1]
		x >>= 1
	}
}

//heapifyDown 从 x 结点开始向子节点，调整结点使得满足堆的性质
//
//注意：你不应该直接使用这个函数
func (pq *PriorityQueue) heapifyDown(x int) {
	for (x << 1) <= len(pq.heap) {
		t := x << 1
		if t+1 <= len(pq.heap) && pq.cmp(pq.heap[t+1-1], pq.heap[t-1]) {
			t++
		}
		if !pq.cmp(pq.heap[t-1], pq.heap[x-1]) {
			break
		}
		pq.heap[x-1], pq.heap[t-1] = pq.heap[t-1], pq.heap[x-1]
		x = t
	}
}

//push 向优先队列中加入一个元素（并自动调整优先队列）
func (pq *PriorityQueue) push(x int) {
	pq.heap = append(pq.heap, x)
	pq.heapifyUp(len(pq.heap))
}

//size 返回优先队列内的元素个数
func (pq *PriorityQueue) size() int { return len(pq.heap) }

//top 返回优先队列中优先级最高的元素
func (pq *PriorityQueue) top() int { return pq.heap[1-1] }

//pop 弹出优先队列中优先级最高的元素
func (pq *PriorityQueue) pop() {
	pq.heap[1-1], pq.heap[len(pq.heap)-1] = pq.heap[len(pq.heap)-1], pq.heap[1-1]
	pq.heap = pq.heap[:len(pq.heap)-1]
	pq.heapifyDown(1)
}

//newPriorityQueue 返回一个指向优先队列的指针
//
//第一个形参：传入一个 int 类型切片，创建优先队列时，你可以用一个 int 类型切片作为初始来创建，也可以设为空切片（nil)
//
//第二个形参：传入一个 func(int, int) bool 类型的函数指针作为比较器，以规定优先队列的优先级。
//传入比较器的两个元素如果满足前者比后者大，则返回 true；否则返回 false。
//
//你可以使用已有的 greater 函数指定小的 int 优先，也可以使用已有的 less 函数指定大的 int 优先
func newPriorityQueue(a []int, fun func(int, int) bool) *PriorityQueue {
	var q PriorityQueue
	copy(q.heap, a)
	q.cmp = fun
	for i := len(a); i >= 1; i-- {
		q.heapifyDown(i)
	}
	return &q
}

//less 判断传入参数 a 和 b 是否满足 a > b
func less(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

//greater 判断传入参数 a 和 b 是否满足 a < b
func greater(a, b int) bool {
	if a < b {
		return true
	} else {
		return false
	}
}
