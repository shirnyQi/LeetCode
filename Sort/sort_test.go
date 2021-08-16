package Sort

import (
	"fmt"
	"sort"
	"testing"
)

// 官方提供的三种排序方式——Ints，Strings，Float64s
func TestCommonFunc(t *testing.T) {
	fmt.Println("====Test Sort Int====")
	a := []int{4, 5, 6, 2, 1, 4, 0}
	fmt.Println(a, sort.IntsAreSorted(a))
	sort.Ints(a)
	fmt.Println(a, sort.IntsAreSorted(a))

	fmt.Println("====Test Sort String====")
	b := []string{"dog", "cat", "mouse", "tiger", "lion"}
	fmt.Println(b, sort.StringsAreSorted(b))
	sort.Strings(b)
	fmt.Println(b, sort.StringsAreSorted(b))

	fmt.Println("====Test Sort Float64====")
	c := []float64{1.23, 4.56, 1.0, 0.56, 0.78}
	fmt.Println(c, sort.Float64sAreSorted(c))
	sort.Float64s(c)
	fmt.Println(c, sort.Float64sAreSorted(c))
}

// 自定义排序方式1——通过回调的方式重载Less函数
func TestOverLoadLessFunc(t *testing.T) {
	fmt.Println("====Test Sort Int====")
	a := []int{4, 5, 6, 2, 1, 4, 0}
	fmt.Println(a, sort.IntsAreSorted(a))

	// 升序
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println(a, sort.IntsAreSorted(a))

	// 降序
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a, sort.SliceIsSorted(a, func(i, j int) bool {
		return a[i] > a[j]
	}))
}

// 自定义排序方式2——重新实现Sort类型接口
// Len, Swap, Less
type student struct {
	name string
	age  int
}

type studentSlice []student

func (s studentSlice) Len() int {
	return len(s)
}

func (s studentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s studentSlice) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func TestSortStudentSliceByAge(t *testing.T) {
	a := studentSlice{
		{
			name: "cat",
			age:  2,
		},
		{
			name: "dog",
			age:  1,
		},
		{
			name: "lion",
			age:  4,
		},
		{
			name: "mouse",
			age:  5,
		},
		{
			name: "tiger",
			age:  2,
		},
	}

	fmt.Println(a, sort.SliceIsSorted(a, a.Less))
	fmt.Println(a, sort.IsSorted(a))
	sort.Sort(a)
	fmt.Println(a, sort.IsSorted(a))
	fmt.Println(a, sort.SliceIsSorted(a, a.Less))
}
