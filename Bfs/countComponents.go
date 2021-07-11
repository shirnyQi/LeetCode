/*
给定编号从 0 到 n-1 的 n 个节点和一个无向边列表（每条边都是一对节点），请编写一个函数来计算无向图中连通分量的数目。

示例1：
输入：
n = 5
edges = [][]int{{0,1},{1,2},{3,4}}

输出：2

示例1：
输入：
n = 5
edges = [][]int{{0,1},{1,2},{2,3},{3,4}}

输出：1

*/

func countComponents(n int, edges [][]int) int {
    graph := make([][]int, n, n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	var (
		res     int = 0
		isVisit     = make([]bool, n, n)
	)

	for i := 0; i < n; i++ {
	if !isVisit[i] {
			bfs(graph, i, isVisit)
			res += 1
		}
	}

	return res

}

type Queue struct {
	queue []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Lenth() int {
	return len(q.queue)
}

func (q *Queue) Enqueue(val int) {
	if q.queue == nil {
		q.queue = make([]int, 0)
	}
	q.queue = append(q.queue, val)
}

func (q *Queue) Dequeue() int {
	if len(q.queue) == 0 {
		return -1
	}
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}

func bfs(graph [][]int, index int, isVisit []bool) {
	queue := NewQueue()
	queue.Enqueue(index)
	isVisit[index] = true
	for queue.Lenth() > 0 {

		node := queue.Dequeue()
		if node < 0 {
			continue
		}

		seqs := graph[node]
		for _, seq := range seqs {
			if !isVisit[seq] {
				isVisit[seq] = true
				queue.Enqueue(seq)
			}
		}
	}
}
