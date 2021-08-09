package dietPlanPerformance

import (
	"fmt"
	"testing"
)

/*
你的好友是一位健身爱好者。前段日子，他给自己制定了一份健身计划。现在想请你帮他评估一下这份计划是否合理。
他会有一份计划消耗的卡路里表，其中 calories[i] 给出了你的这位好友在第 i 天需要消耗的卡路里总量。
为了更好地评估这份计划，对于卡路里表中的每一天，你都需要计算他 「这一天以及之后的连续几天」 （共 k 天）内消耗的总卡路里 T：

如果 T < lower，那么这份计划相对糟糕，并失去 1 分；
如果 T > upper，那么这份计划相对优秀，并获得 1 分；
否则，这份计划普普通通，分值不做变动。
请返回统计完所有 calories.length 天后得到的总分作为评估结果。
注意：总分可能是负数。


示例 1：
	输入：calories = [1,2,3,4,5], k = 1, lower = 3, upper = 3
	输出：0
	解释：calories[0], calories[1] < lower 而 calories[3], calories[4] > upper, 总分 = 0.

示例 2：
	输入：calories = [3,2], k = 2, lower = 0, upper = 1
	输出：1
	解释：calories[0] + calories[1] > upper, 总分 = 1.

示例 3：
	输入：calories = [6,5,0,0], k = 2, lower = 1, upper = 5
	输出：0
	解释：calories[0] + calories[1] > upper, calories[2] + calories[3] < lower, 总分 = 0.

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/ri0jwm/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	var res int

	var left, score int
	for right, v := range calories {
		score += v
		if right-left+1 == k {
			if score > upper {
				res++
			}
			if score < lower {
				res--
			}
			score -= calories[left]
			left++
		}

	}
	return res
}

func TestName(t *testing.T) {
	fmt.Println(dietPlanPerformance([]int{1, 2, 3, 4, 5}, 1, 3, 3))               // 0
	fmt.Println(dietPlanPerformance([]int{3, 2}, 2, 0, 1))                        // 1
	fmt.Println(dietPlanPerformance([]int{6, 5, 0, 0}, 2, 1, 5))                  //0
	fmt.Println(dietPlanPerformance([]int{6, 13, 8, 7, 10, 1, 12, 11}, 6, 5, 37)) //3
}
