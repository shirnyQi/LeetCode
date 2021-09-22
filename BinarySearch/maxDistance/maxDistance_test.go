package maxDistance

import (
	"fmt"
	"sort"
	"testing"
)

/*
1552. 两球之间的磁力
在代号为 C-137 的地球上，Rick 发现如果他将两个球放在他新发明的篮子里，它们之间会形成特殊形式的磁力。Rick 有 n 个空的篮子，第 i 个篮子的位置在 position[i] ，Morty 想把 m 个球放到这些篮子里，使得任意两球间 最小磁力 最大。
已知两个球如果分别位于 x 和 y ，那么它们之间的磁力为 |x - y|
给你一个整数数组 position 和一个整数 m ，请你返回最大化的最小磁力。



示例 1：
输入：position = [1,2,3,4,7], m = 3
输出：3
解释：将 3 个球分别放入位于 1，4 和 7 的三个篮子，两球间的磁力分别为 [3, 3, 6]。最小磁力为 3 。我们没办法让最小磁力大于 3 。

示例 2：
输入：position = [5,4,3,2,1,1000000000], m = 2
输出：999999999
解释：我们使用位于 1 和 1000000000 的篮子时最小磁力最大。

link:https://leetcode-cn.com/problems/magnetic-force-between-two-balls/
*/

func ballsAtMinDis(position []int, minDis int) int {
	n, res, prePos := len(position), 1, 0
	for i := 1; i < n; i++ {
		if position[i]-position[prePos] >= minDis {
			res++
			prePos = i
		}
	}
	return res
}

func maxDistance(position []int, m int) int {
	sort.Ints(position)

	lo, mid, hi := 1, 0, position[len(position)-1]-position[0]+1

	for lo < hi {
		mid = lo + (hi-lo)/2
		balls := ballsAtMinDis(position, mid)
		if balls > m { // mid值估小了，还可以放入更多的球
			lo = mid + 1
		} else if balls == m { // mid值接近了，往右逼近
			lo = mid + 1
		} else { // mid值估多了，球放不下
			hi = mid
		}
	}
	// lo == hi
	return lo - 1
}

func TestName(t *testing.T) {
	fmt.Println(maxDistance([]int{1, 2, 3, 4, 7}, 3))             // 3
	fmt.Println(maxDistance([]int{5, 4, 3, 2, 1, 1000000000}, 2)) // 999999999
}
