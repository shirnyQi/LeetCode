package circularArrayLoop

import (
	"fmt"
	"testing"
)

/*
存在一个不含 0 的 环形 数组 nums ，每个 nums[i] 都表示位于下标 i 的角色应该向前或向后移动的下标个数：
如果 nums[i] 是正数，向前 移动 nums[i] 步
如果 nums[i] 是负数，向后 移动 nums[i] 步
因为数组是 环形 的，所以可以假设从最后一个元素向前移动一步会到达第一个元素，而第一个元素向后移动一步会到达最后一个元素。

数组中的 循环 由长度为 k 的下标序列 seq ：
遵循上述移动规则将导致重复下标序列 seq[0] -> seq[1] -> ... -> seq[k - 1] -> seq[0] -> ...
所有 nums[seq[j]] 应当不是 全正 就是 全负
k > 1
如果 nums 中存在循环，返回 true ；否则，返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/circular-array-loop
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type function interface {
	circularArrayLoop(nums []int) bool
}

type dfs struct {
}

func (b *dfs) circularArrayLoop(nums []int) bool {
	size := len(nums)
	isVisited := make([]bool, size, size)

	var flag bool

	var dfs func(nums []int, index int) bool
	dfs = func(nums []int, index int) bool {
		nextIndex := index + nums[index]%size
		if nextIndex < 0 {
			nextIndex = nextIndex + size
		} else if nextIndex > size-1 {
			nextIndex = nextIndex - size
		}

		if nums[nextIndex] > 0 != flag {
			return false
		}

		if isVisited[nextIndex] {
			if nextIndex != index {
				return true
			} else {
				return false
			}
		}
		isVisited[nextIndex] = true
		if !dfs(nums, nextIndex) {
			return false
		}
		return true
	}

	for i, v := range nums {
		isVisited = make([]bool, size, size)
		flag = v > 0
		if dfs(nums, i) {
			return true
		}
	}

	return false
}

func TestName(t *testing.T) {
	d := &dfs{}
	fmt.Println(d.circularArrayLoop([]int{3, 1, 2}))            //true
	fmt.Println(d.circularArrayLoop([]int{2, -1, 1, 2, 2}))     //true
	fmt.Println(d.circularArrayLoop([]int{-1, 2}))              //false
	fmt.Println(d.circularArrayLoop([]int{-2, 1, -1, -2, -2}))  //false
	fmt.Println(d.circularArrayLoop([]int{-1, -2, -3, -4, -5})) //false
	fmt.Println(d.circularArrayLoop([]int{-2, -3, -9}))         //false
}
