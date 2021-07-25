package findRedundantConnection

import (
	"fmt"
	"testing"
)

/*
树可以看成是一个连通且 无环 的 无向 图。

给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，
且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，
edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。

请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dfs/e3ams1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

var (
	edges1 = [][]int{{1, 2}, {1, 3}, {2, 3}}
	want1  = []int{2, 3}

	edges2 = [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	want2  = []int{1, 5}
)

// 并查集
func findRedundantConnection0(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	fmt.Println(parent)
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}

func findRedundantConnection(edges [][]int) []int {
	var m = make(map[int][]int, 0)
	/*
		0: 未访问
		1: 已访问
		2: 所有子节点已经遍历
	*/
	var isVisited = make([]int, len(edges)+1, len(edges)+1)
	var parent = make([]int, len(edges)+1, len(edges)+1)

	var dfs func(map[int][]int, int) bool
	dfs = func(m map[int][]int, index int) bool {
		isVisited[index] = 1
		for _, next := range m[index] {
			if index == next {
				continue
			}
			if isVisited[next] == 0 {
				parent[next] = index
				dfs(m, next)
			} else if isVisited[next] == 1 && parent[index] != next {
				return true
			}
		}
		isVisited[index] = 2
		return false
	}
	for _, e := range edges {
		if m[e[0]] == nil {
			m[e[0]] = []int{e[1]}
		} else {
			m[e[0]] = append(m[e[0]], e[1])
		}
		if m[e[1]] == nil {
			m[e[1]] = []int{e[0]}
		} else {
			m[e[1]] = append(m[e[1]], e[0])
		}
		if dfs(m, e[0]) {
			return e
		}
		isVisited = make([]int, len(edges)+1, len(edges)+1)
		parent = make([]int, len(edges)+1, len(edges)+1)
	}
	return nil
}

func Test_findRedundantConnection(t *testing.T) {
	//fmt.Println(findRedundantConnection0(edges1))
	fmt.Println(findRedundantConnection(edges2))

}
