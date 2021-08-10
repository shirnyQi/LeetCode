package canReach

import (
	"container/list"
	"testing"
)

/*
来源：力扣（LeetCode）,T1306
链接：https://leetcode-cn.com/problems/jump-game-iii

这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i - arr[i]。
请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。
注意，不管是什么情况下，你都无法跳到数组之外。
*/
var (
	arr1   = []int{4, 2, 3, 0, 3, 1, 2}
	start1 = 0

	arr2   = []int{3, 0, 2, 1, 2}
	start2 = 2

	arr3   = []int{0, 3, 0, 6, 3, 3, 4}
	start3 = 6
)

func canReach(arr []int, start int) bool {
	if arr == nil || start >= len(arr) || start < 0 {
		return false
	}
	var isVisited = make([]bool, len(arr), len(arr))

	q := list.New()
	q.PushBack(start)

	for q.Len() > 0 {
		lenth := q.Len()
		for i := 0; i < lenth; i++ {
			index := q.Remove(q.Front()).(int)
			if arr[index] == 0 {
				return true
			}
			if index-arr[index] >= 0 && !isVisited[index-arr[index]] {
				isVisited[index-arr[index]] = true
				q.PushBack(index - arr[index])
			}

			if index+arr[index] < len(arr) && !isVisited[index+arr[index]] {
				isVisited[index+arr[index]] = true
				q.PushBack(index + arr[index])
			}
		}
	}

	return false
}

func TestCanReach(t *testing.T) {
	if !canReach(arr1, start1) {
		t.Errorf("can reach accrually")
	}

	if canReach(arr2, start2) {
		t.Errorf("can not reach accrually")
	}

	if !canReach(arr3, start3) {
		t.Errorf("can reach accrually")
	}
}
