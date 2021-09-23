package maxAverageRatio

import (
	"container/heap"
	"fmt"
	"testing"
)

/*
1792. 最大平均通过率
一所学校里有一些班级，每个班级里有一些学生，现在每个班都会进行一场期末考试。给你一个二维数组 classes ，其中 classes[i] = [passi, totali] ，表示你提前知道了第 i 个班级总共有 totali 个学生，其中只有 passi 个学生可以通过考试。
给你一个整数 extraStudents ，表示额外有 extraStudents 个聪明的学生，他们 一定 能通过任何班级的期末考。你需要给这 extraStudents 个学生每人都安排一个班级，使得 所有 班级的 平均 通过率 最大 。
一个班级的 通过率 等于这个班级通过考试的学生人数除以这个班级的总人数。平均通过率 是所有班级的通过率之和除以班级数目。
请你返回在安排这 extraStudents 个学生去对应班级后的 最大 平均通过率。与标准答案误差范围在 10-5 以内的结果都会视为正确结果。

示例 1：
输入：classes = [[1,2],[3,5],[2,2]], extraStudents = 2
输出：0.78333
解释：你可以将额外的两个学生都安排到第一个班级，平均通过率为 (3/4 + 3/5 + 2/2) / 3 = 0.78333 。

示例 2：
输入：classes = [[2,4],[3,9],[4,5],[2,10]], extraStudents = 4
输出：0.53485

link:https://leetcode-cn.com/problems/maximum-average-pass-ratio/
*/

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &heapScore{}

	for _, v := range classes {
		heap.Push(h, score{pass: v[0], total: v[1]})
	}

	for extraStudents > 0 {
		s := heap.Pop(h).(score)
		heap.Push(h, score{pass: s.pass + 1, total: s.total + 1})
		extraStudents--
	}

	var res float64

	for h.Len() > 0 {
		s := heap.Pop(h).(score)
		res += float64(s.pass) / float64(s.total)
	}

	return res / float64(len(classes))
}

type score struct {
	pass  int
	total int
}
type heapScore []score

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapScore) Less(i, j int) bool {
	ori1 := float64((*h)[i].pass) / float64((*h)[i].total)
	ori2 := float64((*h)[j].pass) / float64((*h)[j].total)

	aft1 := float64((*h)[i].pass+1) / float64((*h)[i].total+1)
	aft2 := float64((*h)[j].pass+1) / float64((*h)[j].total+1)
	return aft1-ori1 > aft2-ori2
}
func (h *heapScore) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapScore) Len() int      { return len(*h) }
func (h *heapScore) Push(x interface{}) {
	*h = append(*h, x.(score))
}
func (h *heapScore) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func TestName(t *testing.T) {
	fmt.Println(maxAverageRatio([][]int{{1, 2}, {3, 5}, {2, 2}}, 2))
	fmt.Println(maxAverageRatio([][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4))
}
