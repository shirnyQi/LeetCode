package permute

import (
	"container/list"
	"fmt"
	"testing"
)

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dfs/e364om/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// bfs实现
type bfs struct {
}

func (b *bfs) permute(nums []int) [][]int {
	var res [][]int

	var bfs func(nums []int)
	bfs = func(nums []int) {
		queue := list.New()
		queue.PushBack([]int{})

		size := len(nums)
		for queue.Len() > 0 {
			length := queue.Len()
			for i := 0; i < length; i++ {
				path := queue.Remove(queue.Front()).([]int)
				if len(path) == size {
					res = append(res, path)
					continue
				}

				for _, v := range nums {
					isExist := false
					for _, v2 := range path {
						if v2 == v {
							isExist = true
							break
						}
					}

					if isExist {
						continue
					}

					var next []int
					next = append(next, path...)
					next = append(next, v)
					queue.PushBack(next)
				}
			}
		}
	}
	bfs(nums)
	return res
}

// dfs+回溯实现
type dfs struct {
}

func (d *dfs) permute(nums []int) [][]int {
	var res [][]int
	size := len(nums)

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == size {
			res = append(res, path)
			return
		}

		for _, v := range nums {
			isExist := false
			for _, v2 := range path {
				if v2 == v {
					isExist = true
					break
				}
			}

			if isExist {
				continue
			}

			var next []int
			next = append(next, path...)
			next = append(next, v)
			dfs(next)
		}

	}
	dfs([]int{})
	return res
}

func TestName(t *testing.T) {
	b := &bfs{}
	fmt.Println(b.permute([]int{1, 2, 3})) // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
	fmt.Println(b.permute([]int{0, 1}))    // [[0,1],[1,0]]
	fmt.Println(b.permute([]int{1}))       // [[1]]

	d := &dfs{}
	fmt.Println(d.permute([]int{1, 2, 3})) // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
	fmt.Println(d.permute([]int{0, 1}))    // [[0,1],[1,0]]
	fmt.Println(d.permute([]int{1}))       // [[1]]
}
