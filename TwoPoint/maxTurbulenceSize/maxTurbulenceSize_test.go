package maxTurbulenceSize

import (
	"fmt"
	"testing"
)

/*
当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：
若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。
返回 A 的最大湍流子数组的长度。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/rinwtr/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// personal 丑陋的双指针实现，垃圾代码
func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	var res int

	/*
		0：左右相等
		1：左边大于右边
		2：右边大于左边
	*/
	var flag int

	left, right := 0, 1
	if arr[left] > arr[right] {
		flag = 1
	} else if arr[left] < arr[right] {
		flag = 2
	} else {
		flag = 0
	}

	for ; right < len(arr); right++ {
		if flag == 1 && arr[right-1] > arr[right] {
			flag = 2
			continue
		}

		if flag == 2 && arr[right-1] < arr[right] {
			flag = 1
			continue
		}
		fmt.Println(left, right, flag)
		res = max(res, right-left)
		left = right - 1
		if arr[right] > arr[right-1] {
			flag = 1
		} else if arr[right] < arr[right-1] {
			flag = 2
		} else {
			flag = 0
			left++
		}

		fmt.Println(left, right, flag)
	}

	if flag == 1 && arr[len(arr)-1] > arr[len(arr)-2] {
		res = max(res, right-left)
	}
	if flag == 2 && arr[len(arr)-1] < arr[len(arr)-2] {
		res = max(res, right-left)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// standard
func maxTurbulenceSize2(arr []int) int {
	size := len(arr)
	if size < 2 {
		return size
	}

	left, right := 0, 1

	/*
		true:arr[i - 1] < arr[i]
	*/
	var pre bool
	res := 1

	for right < size {
		current := arr[right-1] < arr[right]
		if current == pre {
			left = right - 1
		}
		if arr[right-1] == arr[right] {
			left = right
		}
		right++
		// 不用找到了最终位置之后才去处理左指针和res，需要在移动中知道不变的量
		// 否则就像自己写的方法，一直continue到最终状态，边界值无法判断
		res = max(res, right-left)
		pre = current
	}
	return res
}

func TestName(t *testing.T) {
	fmt.Println(maxTurbulenceSize2([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))           // 5
	fmt.Println(maxTurbulenceSize2([]int{4, 8, 12, 16}))                         // 2
	fmt.Println(maxTurbulenceSize2([]int{100}))                                  //1
	fmt.Println(maxTurbulenceSize2([]int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24})) // 8
	fmt.Println(maxTurbulenceSize2([]int{100, 100, 100}))                        //1

	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))           // 5
	fmt.Println(maxTurbulenceSize([]int{4, 8, 12, 16}))                         // 2
	fmt.Println(maxTurbulenceSize([]int{100}))                                  //1
	fmt.Println(maxTurbulenceSize([]int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24})) // 8
	fmt.Println(maxTurbulenceSize([]int{100, 100, 100}))                        //1
}
