package snakesAndLadders

import (
	"container/list"
	"fmt"
	"testing"
)

/*
909. 蛇梯棋
N x N 的棋盘 board 上，按从 1 到 N*N 的数字给方格编号，编号 从左下角开始，每一行交替方向。
例如，一块 6 x 6 大小的棋盘，编号如下：
r 行 c 列的棋盘，按前述方法编号，棋盘格中可能存在 “蛇” 或 “梯子”；如果 board[r][c] != -1，那个蛇或梯子的目的地将会是 board[r][c]。
玩家从棋盘上的方格 1 （总是在最后一行、第一列）开始出发。
每一回合，玩家需要从当前方格 x 开始出发，按下述要求前进：
选定目标方格：从编号为 x+1，x+2，x+3，x+4，x+5，或者 x+6 的方格中选出一个作为目标方格 s ，目标方格的编号 <= N*N。
该选择模拟了掷骰子的情景，无论棋盘大小如何，你的目的地范围也只能处于区间 [x+1, x+6] 之间。
传送玩家：如果目标方格 S 处存在蛇或梯子，那么玩家会传送到蛇或梯子的目的地。否则，玩家传送到目标方格 S 。
注意，玩家在每回合的前进过程中最多只能爬过蛇或梯子一次：就算目的地是另一条蛇或梯子的起点，你也不会继续移动。
返回达到方格 N*N 所需的最少移动次数，如果不可能，则返回 -1。

示例：
输入：[
[-1,-1,-1,-1,-1,-1],
[-1,-1,-1,-1,-1,-1],
[-1,-1,-1,-1,-1,-1],
[-1,35,-1,-1,13,-1],
[-1,-1,-1,-1,-1,-1],
[-1,15,-1,-1,-1,-1]]
输出：4
解释：
首先，从方格 1 [第 5 行，第 0 列] 开始。
你决定移动到方格 2，并必须爬过梯子移动到到方格 15。
然后你决定移动到方格 17 [第 3 行，第 4 列]，必须爬过蛇到方格 13。
然后你决定移动到方格 14，且必须通过梯子移动到方格 35。
然后你决定移动到方格 36, 游戏结束。
可以证明你需要至少 4 次移动才能到达第 N*N 个方格，所以答案是 4。
*/

func snakesAndLadders(board [][]int) int {
	rows, cols := len(board), len(board[0])
	visit := make([]bool, rows*cols+1, rows*cols+1)
	queue := list.New()
	queue.PushBack(1)

	var res int
	for queue.Len() > 0 {
		length := queue.Len()
		res++
		for i := 0; i < length; i++ {
			value := queue.Remove(queue.Front()).(int)
			if value == rows*cols {
				return res - 1
			}
			visit[value] = true
			for j := 1; j < 7 && value+j <= rows*cols; j++ {
				var next int
				x, y := getIndex(value+j, rows, cols)
				if board[x][y] == -1 {
					next = value + j
				} else {
					next = board[x][y]
				}
				if !visit[next] {
					queue.PushBack(next)
				}
			}
		}
	}

	return -1
}

func getIndex(value, rows, cols int) (int, int) {
	var x, y int
	line := (value - 1) / rows
	x = rows - 1 - line

	if line%2 == 0 {
		y = value - 1 - line*rows
	} else {
		y = cols - (value - line*rows)
	}
	return x, y
}

func TestName(t *testing.T) {
	var board [][]int
	board = [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1}}

	fmt.Println(snakesAndLadders(board)) // 4

	board = [][]int{
		{-1, -1, 27, 13, -1, 25, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{44, -1, 8, -1, -1, 2, -1},
		{-1, 30, -1, -1, -1, -1, -1},
		{3, -1, 20, -1, 46, 6, -1},
		{-1, -1, -1, -1, -1, -1, 29},
		{-1, 29, 21, 33, -1, -1, -1}}
	fmt.Println(snakesAndLadders(board)) // 4
}
