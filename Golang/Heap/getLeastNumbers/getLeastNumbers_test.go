package getLeastNumbers

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]

作者：爱学习的饲养员
链接：https://leetcode-cn.com/leetbook/read/heap/ev9bph/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func getLeastNumbersSort(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

type heapInt []int

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapInt) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
func (h *heapInt) Peek() int {
	return (*h)[0]
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}

	d := &heapInt{}

	for _, v := range arr {
		if d.Len() < k {
			heap.Push(d, v)
		} else {
			if d.Peek() > v {
				heap.Pop(d)
				heap.Push(d, v)
			}
		}
	}

	return *d
}

func TestName(t *testing.T) {
	fmt.Println(getLeastNumbersSort([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbersSort([]int{0, 1, 2, 1}, 1))

	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 1))
}
