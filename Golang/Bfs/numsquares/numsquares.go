package numsquares

import "container/list"

/*
LC T279:
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
*/

func numSquares(n int) int {
	queue := list.New()
	isVisited := make([]bool, n+1, n+1)

	queue.PushBack(n)

	var step int = 1
	for queue.Len() > 0 {
		lenth := queue.Len()
		for i := 0; i < lenth; i++ {
			number := queue.Remove(queue.Front()).(int)
			for j := 0; j*j <= number; j++ {
				if j*j == number {
					return step
				}
				next := number - j*j
				if !isVisited[next] {
					queue.PushBack(next)
					isVisited[next] = true
				}
			}
		}
		step++
	}
	return 0
}
