/*
lc:https://leetcode-cn.com/problems/max-consecutive-ones-iii/
*/

var (
	A1 = []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	K1 = 2

	A2 = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	K2 = 3

	A3 = []int{1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1}
	K3 = 9
)

func longestOnes(A []int, K int) int {
	var left int
	count := K
	maxLength := 0
	for right, val := range A {
		if val == 0 {
			count--
		}
		for count < 0 {
			if A[left] == 0 {
				count++
			}
			left++
		}
		fmt.Println(left, right)
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
