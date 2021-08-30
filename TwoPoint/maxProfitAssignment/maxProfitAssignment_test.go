package maxProfitAssignment

import (
	"fmt"
	"sort"
	"testing"
)

/*安排工作以达到最大收益
有一些工作：difficulty[i] 表示第 i 个工作的难度，profit[i] 表示第 i 个工作的收益。
现在我们有一些工人。worker[i] 是第 i 个工人的能力，即该工人只能完成难度小于等于 worker[i] 的工作。
每一个工人都最多只能安排一个工作，但是一个工作可以完成多次。
举个例子，如果 3 个工人都尝试完成一份报酬为 1 的同样工作，那么总收益为 $3。如果一个工人不能完成任何工作，他的收益为 $0 。

我们能得到的最大收益是多少？

示例：

输入: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
输出: 100
解释: 工人被分配的工作难度是 [4,4,6,6] ，分别获得 [20,20,30,30] 的收益。
*/

type work struct {
	diff   int
	profit int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	size := len(difficulty)

	works := make([]work, size, size)

	for i := 0; i < size; i++ {
		works[i] = work{diff: difficulty[i], profit: profit[i]}
	}

	sort.Slice(works, func(i, j int) bool {
		if works[i].diff < works[j].diff {
			return true
		} else if works[i].diff > works[j].diff {
			return false
		} else {
			return works[i].profit < works[j].profit
		}
	})

	sort.Ints(worker)

	fmt.Println(works)

	var res int

	i, j, t := 0, 0, 0

	for i < len(worker) {
		for j < size && worker[i] >= works[j].diff {
			t = max(t, works[j].profit)
			j++
		}
		res += t
		i++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))           // 100
	fmt.Println(maxProfitAssignment([]int{68, 35, 52, 47, 86}, []int{67, 17, 1, 81, 3}, []int{92, 10, 85, 84, 82})) // 324
}
