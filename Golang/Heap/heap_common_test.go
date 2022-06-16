package Heap

import (
	"container/heap"
	"fmt"
	"testing"
)

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

func TestHeap(t *testing.T) {
	h := &heapInt{}

	heap.Push(h, 1)
	heap.Push(h, 2)
	heap.Push(h, 3)
	fmt.Println(h)        // 大根堆，3，1，2
	fmt.Println(h.Peek()) // 堆顶为3

	heap.Pop(h) // 2,1
	fmt.Println(h)

	heap.Push(h, 4)
	heap.Push(h, 0)
	fmt.Println(h) // 4,1,2,0
}
