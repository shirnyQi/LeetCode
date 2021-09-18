package asteroidCollision

import (
	"fmt"
	"testing"
)

/*
735. 行星碰撞
给定一个整数数组 asteroids，表示在同一行的行星。
对于数组中的每一个元素，其绝对值表示行星的大小，正负表示行星的移动方向（正表示向右移动，负表示向左移动）。每一颗行星以相同的速度移动。
找出碰撞后剩下的所有行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。


示例 1：
输入：asteroids = [5,10,-5]
输出：[5,10]
解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。

示例 2：
输入：asteroids = [8,-8]
输出：[]
解释：8 和 -8 碰撞后，两者都发生爆炸。

示例 3：
输入：asteroids = [10,2,-5]
输出：[10]
解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。

示例 4：
输入：asteroids = [-2,-1,1,2]
输出：[-2,-1,1,2]
解释：-2 和 -1 向左移动，而 1 和 2 向右移动。 由于移动方向相同的行星不会发生碰撞，所以最终没有行星发生碰撞。

link:https://leetcode-cn.com/problems/asteroid-collision/
*/

func asteroidCollision(asteroids []int) []int {
	var res []int

	for _, v := range asteroids {
		if len(res) == 0 {
			res = append(res, v)
			continue
		}

		top := res[len(res)-1]
		if noColl(top, v) {
			res = append(res, v)
			continue
		}
		for top+v <= 0 {
			if top+v == 0 {
				res = res[:len(res)-1]
				break
			}
			if noColl(top, v) || top == 0 {
				res = append(res, v)
				break
			}

			res = res[:len(res)-1]
			if len(res) > 0 {
				top = res[len(res)-1]
			} else {
				top = 0
			}
		}
	}
	return res
}

func noColl(top, v int) bool {
	if (top > 0 && v > 0) || (top < 0 && v < 0) || (top < 0 && v > 0) {
		return true
	}
	return false
}

func TestName(t *testing.T) {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2}))
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -1}))
	fmt.Println(asteroidCollision([]int{1, -2, -2, -2}))
}
