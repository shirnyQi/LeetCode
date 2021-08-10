package canMeasureWater

import (
	"testing"
)

/*
作者：力扣 (LeetCode) T365
链接：https://leetcode-cn.com/leetbook/read/bfs/e6fp6m/

有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
你允许：
装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
*/

var (
	x1, y1, z1 = 3, 5, 4
	x2, y2, z2 = 2, 6, 5
)

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity == targetCapacity || jug2Capacity == targetCapacity || targetCapacity == 0 {
		return true
	}

	if jug1Capacity+jug2Capacity < targetCapacity {
		return false
	}

	return bfs(jug1Capacity, jug2Capacity, targetCapacity)
}

type State struct {
	X int
	Y int
}

func bfs(x, y, z int) bool {
	var queue []State
	initState := State{
		X: 0,
		Y: 0,
	}
	queue = append(queue, initState)
	visitedMap := make(map[State]bool)

	for len(queue) > 0 {
		nowState := queue[0]
		queue = queue[1:]
		if visitedMap[nowState] {
			continue
		}
		visitedMap[nowState] = true

		// 截止条件
		if nowState.X == z || nowState.Y == z || nowState.X+nowState.Y == z {
			return true
		}

		// 增加下一次可达的状态：倒满X
		nextState := State{
			X: x,
			Y: nowState.Y,
		}
		queue = append(queue, nextState)

		// 增加下一次可达的状态：倒满Y
		nextState = State{
			X: nowState.X,
			Y: y,
		}
		queue = append(queue, nextState)

		// 增加下一次可达的状态：清空X
		nextState = State{
			X: 0,
			Y: nowState.Y,
		}
		queue = append(queue, nextState)

		// 增加下一次可达的状态：清空Y
		nextState = State{
			X: nowState.X,
			Y: 0,
		}
		queue = append(queue, nextState)

		// 增加下一次可达的状态：X的水倒入Y
		if nowState.X > y-nowState.Y {
			nextState = State{
				X: nowState.X + nowState.Y - y,
				Y: y,
			}
			queue = append(queue, nextState)
		} else {
			nextState = State{
				X: 0,
				Y: nowState.X + nowState.Y,
			}
			queue = append(queue, nextState)
		}

		// 增加下一次可达的状态：Y的水倒入X
		if nowState.Y > x-nowState.X {
			nextState = State{
				X: x,
				Y: nowState.X + nowState.Y - x,
			}
			queue = append(queue, nextState)
		} else {
			nextState = State{
				X: nowState.X + nowState.Y,
				Y: 0,
			}
			queue = append(queue, nextState)
		}
	}
	return false
}

func TestCanMeasureWater(t *testing.T) {
	if !canMeasureWater(x1, y1, z1) {
		t.Errorf("can measure accrually")
	}

	if canMeasureWater(x2, y2, z2) {
		t.Errorf("can not measure accrually")
	}
}
