package eventualSafeNodes

import (
	"fmt"
	"testing"
)

/*
在有向图中，从某个节点和每个转向处开始出发，沿着图的有向边走。如果到达的节点是终点（即它没有连出的有向边），则停止。

如果从起始节点出发，最后必然能走到终点，就认为起始节点是 最终安全 的。更具体地说，
对于最终安全的起始节点而言，存在一个自然数 k ，无论选择沿哪条有向边行走 ，走了不到 k 步后必能停止在一个终点上。

返回一个由图中所有最终安全的起始节点组成的数组作为答案。答案数组中的元素应当按 升序 排列。

该有向图有 n 个节点，按 0 到 n - 1 编号，其中 n 是 graph 的节点数。图以下述形式给出：graph[i] 是编号 j 节点的一个列表，满足 (i, j) 是图的一条有向边。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dfs/e32oz7/
*/

var (
	graph1 = [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	want1  = []int{2, 4, 5, 6}

	graph2 = [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}
	want2  = []int{4}
)

func eventualSafeNodes(graph [][]int) []int {
	var res []int

	var size = len(graph)

	/*
		0: 未访问
		1: 已访问
		2: 所有子节点已经遍历
	*/
	var isVisited = make([]int, size, size)

	var dfs func([][]int, []int, int) bool
	dfs = func(graph [][]int, isVisited []int, index int) bool {
		isVisited[index] = 1

		for _, next := range graph[index] {
			if isVisited[next] == 0 {
				// 未访问的邻居节点，进行DFS
				if !dfs(graph, isVisited, next) {
					return false
				}
			} else if isVisited[next] == 1 {
				// 已经访问的，说明已经跳转成环，则不是安全点
				return false
			}

		}
		// 全部访问完当前节点的邻居节点，没有遇到成环的情况，说明安全
		isVisited[index] = 2
		return true
	}

	for i := 0; i < size; i++ {
		if dfs(graph, isVisited, i) {
			res = append(res, i)
		}
	}

	return res
}

func Test_eventualSafeNodes(t *testing.T) {
	fmt.Println(eventualSafeNodes(graph1), want1)
	fmt.Println(eventualSafeNodes(graph2), want2)
}
