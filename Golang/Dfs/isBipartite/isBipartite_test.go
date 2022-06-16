package isBipartite

import (
	"container/list"
	"fmt"
	"testing"
)

/*
存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。
给你一个二维数组 graph ，其中 graph[u] 是一个节点数组，由节点 u 的邻接节点组成。
形式上，对于 graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。该无向图同时具有以下属性：
不存在自环（graph[u] 不包含 u）。
不存在平行边（graph[u] 不包含重复值）。
如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
二分图 定义：如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，
并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 二分图 。

如果图是二分图，返回 true ；否则，返回 false 。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dfs/e3n5g5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

var (
	graph1 = [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	graph2 = [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
)

// DFS
func isBipartite(graph [][]int) bool {
	var size = len(graph)
	var isValid = true

	var isVisited = make([]int, size, size)

	var dfs func([][]int, int, int)
	dfs = func(graph [][]int, state, index int) {
		fmt.Println(state, index)
		isVisited[index] = state
		nextState := 1
		if state == 1 {
			nextState = 2
		}

		for _, next := range graph[index] {
			if isVisited[next] == 0 {
				dfs(graph, nextState, next)
				if !isValid {
					return
				}
			} else if isVisited[next] != nextState {
				isValid = false
				return
			}
		}

		return
	}

	for i := 0; i < size && isValid; i++ {
		if isVisited[i] == 0 {
			dfs(graph, 1, i)
		}
	}

	return isValid
}

// BFS
func isBipartiteBFS(graph [][]int) bool {
	var size = len(graph)
	var isVisited = make([]int, size, size)

	for i := 0; i < size; i++ {
		if isVisited[i] == 0 {
			queue := list.New()
			queue.PushBack(i)
			isVisited[i] = 1

			for queue.Len() > 0 {
				lenth := queue.Len()
				for j := 0; j < lenth; j++ {
					crt := queue.Remove(queue.Front()).(int)
					nextState := 1
					if isVisited[crt] == 1 {
						nextState = 2
					}
					for _, next := range graph[crt] {
						if isVisited[next] == 0 {
							isVisited[next] = nextState
							queue.PushBack(next)
						} else if isVisited[next] != nextState {
							return false
						}
					}

				}

			}
		}
	}
	return true
}

func Test_isBipartite(t *testing.T) {
	if isBipartite(graph1) != false {
		t.Errorf("error")
	}

	if isBipartite(graph2) != true {
		t.Errorf("error")
	}

	if isBipartiteBFS(graph1) != false {
		t.Errorf("error")
	}

	if isBipartiteBFS(graph2) != true {
		t.Errorf("error")
	}
}
