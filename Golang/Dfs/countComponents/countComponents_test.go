package countComponents

func countComponents(n int, edges [][]int) int {
	graph := make([][]int, n, n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	var (
		res     = 0
		isVisit = make([]bool, n, n)
	)

	var dfs func([][]int, []bool, int)
	dfs = func(graph [][]int, isVisit []bool, index int) {
		isVisit[index] = true

		for _, x := range graph[index] {
			if !isVisit[x] {
				dfs(graph, isVisit, x)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !isVisit[i] {
			res++
			dfs(graph, isVisit, i)
		}
	}

	return res
}
