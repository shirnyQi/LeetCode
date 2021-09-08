package KthLargest

import (
	"container/heap"
	"fmt"
	"testing"
)

/*
703. 数据流中的第 K 大元素
设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

请实现 KthLargest 类：

KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。


示例：

输入：
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
输出：
[null, 4, 5, 5, 8, 8]

解释：
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8

链接：https://leetcode-cn.com/leetbook/read/heap/evozem/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

type KthLargest struct {
	kth  int
	heap *heapInt
}

func Constructor(k int, nums []int) KthLargest {
	var h = &heapInt{}

	for _, v := range nums {
		if h.Len() < k {
			heap.Push(h, v)
		} else {
			if h.Peek() < v {
				heap.Pop(h)
				heap.Push(h, v)
			}
		}
	}

	return KthLargest{kth: k, heap: h}
}

func (this *KthLargest) Add(val int) int {
	if this.heap.Len() < this.kth {
		heap.Push(this.heap, val)
		return this.heap.Peek()
	}

	if this.heap.Peek() < val {
		heap.Pop(this.heap)
		heap.Push(this.heap, val)
	}

	return this.heap.Peek()
}

type heapInt []int

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
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

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func TestName(t *testing.T) {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))  // 4
	fmt.Println(obj.Add(5))  // 5
	fmt.Println(obj.Add(10)) // 5
	fmt.Println(obj.Add(9))  // 8
	fmt.Println(obj.Add(4))  // 8
}
