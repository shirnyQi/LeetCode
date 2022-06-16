package maximumUniqueSubarray

import (
	"fmt"
	"testing"
)

// standard
func maximumUniqueSubarray(nums []int) int {
	m := make(map[int]int, 0)

	left, right := 0, 0
	maxSum, tmpSum := 0, 0

	for ; right < len(nums); right++ {
		tmpSum += nums[right]
		m[nums[right]]++
		for m[nums[right]] > 1 {
			tmpSum -= nums[left]
			m[nums[left]]--
			left++
		}
		maxSum = max(maxSum, tmpSum)
	}
	return maxSum
}


// personal
func maximumUniqueSubarray2(nums []int) int {
	m := make(map[int]int, 0)

	left, right := 0, 0
	maxSum, tmpSum := 0, 0

	nums = append(nums, 0)

	for right < len(nums) {
		if m[nums[right]] == 0 && nums[right] != 0 {
			tmpSum += nums[right]
			m[nums[right]] = 1
		} else {
			maxSum = max(maxSum, tmpSum)
			for nums[left] != nums[right] && left <= right {
				tmpSum -= nums[left]
				m[nums[left]] = 0
				left++
			}
			left++
		}
		right++
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))             //17
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5})) //8
	fmt.Println(maximumUniqueSubarray([]int{
		187, 470, 25, 436, 538, 809, 441, 167, 477,
		110, 275, 133, 666, 345, 411, 459, 490, 266, 987, 965, 429,
		166, 809, 340, 467, 318, 125, 165, 809, 610, 31, 585, 970,
		306, 42, 189, 169, 743, 78, 810, 70, 382, 367, 490, 787, 670,
		476, 278, 775, 673, 299, 19, 893, 817, 971, 458, 409, 886, 434})) //16911

	fmt.Println(maximumUniqueSubarray2([]int{4, 2, 4, 5, 6}))             //17
	fmt.Println(maximumUniqueSubarray2([]int{5, 2, 1, 2, 5, 2, 1, 2, 5})) //8
	fmt.Println(maximumUniqueSubarray2([]int{
		187, 470, 25, 436, 538, 809, 441, 167, 477,
		110, 275, 133, 666, 345, 411, 459, 490, 266, 987, 965, 429,
		166, 809, 340, 467, 318, 125, 165, 809, 610, 31, 585, 970,
		306, 42, 189, 169, 743, 78, 810, 70, 382, 367, 490, 787, 670,
		476, 278, 775, 673, 299, 19, 893, 817, 971, 458, 409, 886, 434})) //16911
}

