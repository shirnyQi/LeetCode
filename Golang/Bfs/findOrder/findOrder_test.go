package findOrder

import (
	"container/list"
	"fmt"
	"testing"
)

/*
剑指 Offer II 113. 课程顺序
现在总共有 numCourses 门课需要选，记为 0 到 numCourses-1。
给定一个数组 prerequisites ，它的每一个元素 prerequisites[i] 表示两门课程之间的先修顺序。 例如 prerequisites[i] = [ai, bi] 表示想要学习课程 ai ，需要先完成课程 bi 。
请根据给出的总课程数  numCourses 和表示先修顺序的 prerequisites 得出一个可行的修课序列。
可能会有多个正确的顺序，只要任意返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

示例 1:
输入: numCourses = 2, prerequisites = [[1,0]]
输出: [0,1]
解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。

示例 2:
输入: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出: [0,1,2,3] or [0,2,1,3]
解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。

 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
示例 3:
输入: numCourses = 1, prerequisites = []
输出: [0]
解释: 总共 1 门课，直接修第一门课就可。

link:https://leetcode-cn.com/problems/QA2IGt/
*/

// dfs
func findOrderDfs(numCourses int, prerequisites [][]int) []int {
	m := make(map[int][]int, 0)
	// 0:未选修;1:完成中; 2:已完成
	isStudy := make([]int, numCourses, numCourses)

	for _, v := range prerequisites {
		m[v[0]] = append(m[v[0]], v[1])
	}
	var res []int
	var dfs func(course int)
	dfs = func(course int) {
		if isStudy[course] == 2 {
			return
		}
		requires, ok := m[course]
		if !ok {
			res = append(res, course)
			isStudy[course] = 2
			return
		}

		isStudy[course] = 1
		for _, c := range requires {
			if isStudy[c] == 1 {
				return
			}
			dfs(c)
		}
		res = append(res, course)
		isStudy[course] = 2
	}

	for i := 0; i < numCourses; i++ {
		if isStudy[i] == 0 {
			dfs(i)
		}
	}

	if len(res) != numCourses {
		return nil
	}
	return res
}

// bfs
func findOrder(numCourses int, prerequisites [][]int) []int {
	m := make(map[int][]int, 0)
	degree := make([]int, numCourses, numCourses)
	for _, v := range prerequisites {
		m[v[1]] = append(m[v[1]], v[0])
		degree[v[0]]++
	}

	queue := list.New()
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			queue.PushBack(i)
		}
	}

	var res []int

	for queue.Len() > 0 {
		course := queue.Remove(queue.Front()).(int)
		res = append(res, course)

		for _, c := range m[course] {
			degree[c]--
			if degree[c] == 0 {
				queue.PushBack(c)
			}
		}
	}

	if len(res) != numCourses {
		return nil
	}
	return res
}

func TestName(t *testing.T) {
	fmt.Println(findOrderDfs(2, [][]int{{1, 0}}))                         //[0,1]
	fmt.Println(findOrderDfs(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})) // [0,1,2,3] or [0,2,1,3]
	fmt.Println(findOrderDfs(1, [][]int{}))                               // [0]
	fmt.Println(findOrderDfs(2, [][]int{{1, 0}, {0, 1}}))                 //[0,1]

	fmt.Println(findOrder(2, [][]int{{1, 0}}))                         //[0,1]
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})) // [0,1,2,3] or [0,2,1,3]
	fmt.Println(findOrder(1, [][]int{}))                               // [0]
	fmt.Println(findOrder(2, [][]int{{1, 0}, {0, 1}}))                 //[0,1]
}
