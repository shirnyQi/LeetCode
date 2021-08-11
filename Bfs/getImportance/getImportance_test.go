package getImportance

import (
	"container/list"
	"testing"
)

/*
作者：力扣 (LeetCode) T690
链接：https://leetcode-cn.com/leetbook/read/bfs/e6bwzh/
给定一个保存员工信息的数据结构，它包含了员工 唯一的 id ，重要度 和 直系下属的 id 。

比如，员工 1 是员工 2 的领导，员工 2 是员工 3 的领导。他们相应的重要度为 15 , 10 , 5 。
那么员工 1 的数据结构是 [1, 15, [2]] ，员工 2的 数据结构是 [2, 10, [3]] ，员工 3 的数据结构是 [3, 5, []] 。
注意虽然员工 3 也是员工 1 的一个下属，但是由于 并不是直系 下属，因此没有体现在员工 1 的数据结构中。

现在输入一个公司的所有员工信息，以及单个员工 id ，返回这个员工和他所有下属的重要度之和。
*/

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

var (
	E1      = &Employee{1, 5, []int{2, 3}}
	E2      = &Employee{2, 3, []int{}}
	E3      = &Employee{3, 3, []int{}}
	startId = 1

	want = 11 // 5 + 3 + 3
)

func getImportance(employees []*Employee, id int) int {
	employeeMap := make(map[int]*Employee, 0)
	for i := range employees {
		employeeMap[employees[i].Id] = employees[i]
	}
	if employeeMap[id] == nil {
		return 0
	}
	return Bfs(employeeMap, employeeMap[id])
}

func Bfs(employees map[int]*Employee, employee *Employee) int {
	var result int

	q := list.New()
	q.PushBack(employee)

	for q.Len() > 0 {
		length := q.Len()
		for i := 0; i < length; i++ {
			e := q.Remove(q.Front()).(*Employee)
			result += e.Importance

			for _, subIdx := range e.Subordinates {
				subE := employees[subIdx]
				if subE != nil {
					q.PushBack(subE)
				}
			}
		}
	}
	return result
}

func TestGetImportance(t *testing.T) {
	input := []*Employee{E1, E2, E3}
	var got int
	got = getImportance(input, startId)
	if want != got {
		t.Errorf("err,want: %v,got %v", want, got)
	}
}
