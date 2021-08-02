package characterReplacement

import (
	"fmt"
	"testing"
)

func characterReplacement(s string, k int) int {
	lenth := len(s)
	if lenth < 2 {
		return lenth
	}

	cntMap := make(map[byte]int, 0)

	left, right := 0, 0
	res, maxLenth := 0, 0

	for right < lenth {
		cntMap[[]byte(s)[right]]++
		maxLenth = max(maxLenth, cntMap[[]byte(s)[right]])
		right++
		fmt.Println(cntMap, left, right, maxLenth, k)

		if right-left > maxLenth+k {
			cntMap[[]byte(s)[left]]--
			left++
		}
		fmt.Println(cntMap, left, right, maxLenth, k)
		fmt.Println("========")
		res = max(res, right-left)
	}

	return res
}

func characterReplacement2(s string, k int) int {
	lenth := len(s)
	if lenth < 2 {
		return lenth
	}

	charArray := []byte(s)
	freq := make([]int, 26, 26)

	left, right := 0, 0
	res, maxLenth := 0, 0

	for right < lenth {
		freq[charArray[right]-'A']++
		maxLenth = max(maxLenth, freq[charArray[right]-'A'])
		right++

		if (right - left) > (maxLenth + k) {
			freq[charArray[left]-'A']--
			left++
		}

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

func Test_characterReplacement(t *testing.T) {
	fmt.Println(characterReplacement2("AAABBBCCC", 2)) // 5
	fmt.Println(characterReplacement2("ABAB", 2))      // 4
	fmt.Println(characterReplacement2("AABABBA", 1))   // 4
}
