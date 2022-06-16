package pathInZigZagTree

import (
	"fmt"
	"testing"
)

/*
在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。
如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；
而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。

给你树上某一个节点的标号 label，请你返回从根节点到该标号为 label 节点的路径，该路径是由途经的节点标号所组成的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func pathInZigZagTree(label int) []int {
	var crtNums, upIndex = 1, 0
	for crtNums <= label {
		crtNums *= 2
	}

	var res []int
	var findx func(label int, crtNums int)

	findx = func(label int, crtNums int) {
		if label == 1 {
			res = append([]int{1}, res...)
			return
		}

		res = append([]int{label}, res...)
		upNums := crtNums / 2
		if (crtNums*2-label)%2 == 0 {
			upIndex = (crtNums*2 - label) / 2
		} else {
			upIndex = ((crtNums*2 - label) / 2) + 1
		}
		findx(crtNums-upNums-1+upIndex, upNums)
	}

	findx(label, crtNums/2)
	return res
}

func Test_(t *testing.T) {
	fmt.Println(pathInZigZagTree(1))  // 1
	fmt.Println(pathInZigZagTree(14)) // 1,3,4,14
	fmt.Println(pathInZigZagTree(26)) // 1,2,6,10,26
	fmt.Println(pathInZigZagTree(16)) // 1,3,4,15,16
}
