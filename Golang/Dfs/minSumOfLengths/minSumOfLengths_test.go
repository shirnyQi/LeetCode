package minSumOfLengths

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
1477. 找两个和为目标值且不重叠的子数组
给你一个整数数组 arr 和一个整数值 target 。
请你在 arr 中找 两个互不重叠的子数组 且它们的和都等于 target 。可能会有多种方案，请你返回满足要求的两个子数组长度和的 最小值 。
请返回满足要求的最小长度和，如果无法找到这样的两个子数组，请返回 -1 。

示例 1：
输入：arr = [3,2,2,4,3], target = 3
输出：2
解释：只有两个子数组和为 3 （[3] 和 [3]）。它们的长度和为 2 。

示例 2：
输入：arr = [7,3,4,7], target = 7
输出：2
解释：尽管我们有 3 个互不重叠的子数组和为 7 （[7], [3,4] 和 [7]），但我们会选择第一个和第三个子数组，因为它们的长度和 2 是最小值。

示例 3：
输入：arr = [4,3,2,6,2,3,4], target = 6
输出：-1
解释：我们只有一个和为 6 的子数组。

示例 4：
输入：arr = [5,5,4,4,5], target = 3
输出：-1
解释：我们无法找到和为 3 的子数组。

示例 5：
输入：arr = [3,1,1,1,5,1,2,1], target = 3
输出：3
解释：注意子数组 [1,2] 和 [2,1] 不能成为一个方案因为它们重叠了。
*/

// error
func minSumOfLengthserror(arr []int, target int) int {
	sort.Ints(arr)

	n := len(arr)

	var h = &heapInt{}

	var dfs func(index int, target int, cnt int) int
	dfs = func(index int, target int, cnt int) int {
		if target == 0 {
			return cnt
		}

		if index <= 0 || target < 0 {
			return -1
		}

		return dfs(index-1, target-arr[index], cnt+1)
	}

	for i := n - 1; i >= 0; i-- {
		if arr[i] > target {
			continue
		}

		size := dfs(i, target-arr[i], 1)
		if size == -1 {
			continue
		}
		if h.Len() < 2 {
			heap.Push(h, size)
		} else {
			heap.Pop(h)
			heap.Push(h, size)
		}
	}

	if h.Len() < 2 {
		return -1
	}

	return heap.Pop(h).(int) + heap.Pop(h).(int)
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

// dp
func minSumOfLengths(arr []int, target int) int {
	var leftMin, rightMin = make([]int, len(arr)), make([]int, len(arr))
	var i, j, sum, ans int
	for j = 0; j < len(arr); j++ {
		sum += arr[j]
		for sum > target {
			sum -= arr[i]
			i++
		}
		if j == 0 {
			leftMin[j] = math.MaxInt32
		} else {
			leftMin[j] = leftMin[j-1]
		}
		if sum == target {
			diff := j - i + 1
			if leftMin[j] > diff {
				leftMin[j] = diff
			}
		}
	}
	j = len(arr) - 1
	sum = 0
	for i = len(arr) - 1; i >= 0; i-- {
		sum += arr[i]
		for sum > target {
			sum -= arr[j]
			j--
		}
		if i == len(arr)-1 {
			rightMin[i] = math.MaxInt32
		} else {
			rightMin[i] = rightMin[i+1]
		}
		if sum == target {
			diff := j - i + 1
			if rightMin[i] > diff {
				rightMin[i] = diff
			}
		}

	}
	ans = math.MaxInt32
	for i = 0; i < len(arr)-1; i++ {
		if leftMin[i] == math.MaxInt32 || rightMin[i+1] == math.MaxInt32 {
			continue
		}
		if leftMin[i]+rightMin[i+1] < ans {
			ans = leftMin[i] + rightMin[i+1]
		}
	}
	if ans == math.MaxInt32 {
		return -1
	} else {
		return ans
	}
}

func TestName(t *testing.T) {
	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3))          // 2
	fmt.Println(minSumOfLengths([]int{7, 3, 4, 7}, 7))             // 2
	fmt.Println(minSumOfLengths([]int{4, 3, 2, 6, 2, 3, 4}, 6))    // -1
	fmt.Println(minSumOfLengths([]int{5, 5, 4, 4, 5}, 3))          // -1
	fmt.Println(minSumOfLengths([]int{3, 1, 1, 1, 5, 1, 2, 1}, 3)) // 3
}
