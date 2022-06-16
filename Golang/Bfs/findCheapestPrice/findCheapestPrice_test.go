package findCheapestPrice

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

/*
K 站中转内最便宜的航班
有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城市 fromi 开始，以价格 pricei 抵达 toi。
现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。

示例 1：
输入:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
输出: 200
解释:
城市航班图如下
从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。

示例 2：
输入:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
输出: 500
解释:
城市航班图如下

从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。

link:https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/
*/

type node struct {
	station int
	price   int
}

// BFS
func findCheapestPriceBfs(n int, flights [][]int, src int, dst int, k int) int {
	priceMap := make(map[int][]*node, 0)

	for _, v := range flights {
		priceMap[v[0]] = append(priceMap[v[0]], &node{v[1], v[2]})
	}

	var res = math.MaxInt32

	srcNode := &node{src, 0}
	queue := list.New()
	queue.PushBack(srcNode)

	var layer = -2
	for queue.Len() > 0 {
		layer++
		length := queue.Len()
		for i := 0; i < length; i++ {
			crtNode := queue.Remove(queue.Front()).(*node)
			if crtNode.station == dst && layer <= k {
				res = min(res, crtNode.price)
				continue
			}
			if layer > k {
				continue
			}

			for _, n := range priceMap[crtNode.station] {
				//n.price += crtNode.price
				queue.PushBack(&node{n.station, n.price + crtNode.price})
			}
		}
	}

	if res == math.MaxInt32 {
		return -1
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// dfs
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	priceMap := make(map[int][]*node, 0)

	for _, v := range flights {
		priceMap[v[0]] = append(priceMap[v[0]], &node{v[1], v[2]})
	}
	var isVisited = make([]bool, n+1, n+1)

	var res = math.MaxInt32

	var dfs func(station int, layer int, price int)

	dfs = func(station int, layer int, price int) {
		if station == dst && layer <= k {
			res = min(res, price)
			return
		}

		if layer > k {
			return
		}

		layer++

		for _, n := range priceMap[station] {
			if !isVisited[n.station] {
				isVisited[n.station] = true
				dfs(n.station, layer, n.price+price)
				isVisited[n.station] = false
			}

		}
	}

	isVisited[src] = true
	dfs(src, -1, 0)

	if res == math.MaxInt32 {
		return -1
	}

	return res
}

// DP
func findCheapestPriceDP(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1
	f := make([][]int, k+2)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][src] = 0
	for t := 1; t <= k+1; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			f[t][i] = min(f[t][i], f[t-1][j]+cost)
		}
	}
	ans := inf
	for t := 1; t <= k+1; t++ {
		ans = min(ans, f[t][dst])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

func TestName(t *testing.T) {
	fmt.Println(findCheapestPriceBfs(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)) // 200
	fmt.Println(findCheapestPriceBfs(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0)) // 500

	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)) // 200
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0)) // 500

	fmt.Println(findCheapestPriceDP(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)) // 200
	fmt.Println(findCheapestPriceDP(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0)) // 500
}
