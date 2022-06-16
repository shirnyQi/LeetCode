package FreqStack

import (
	"fmt"
	"testing"
)

/*
895. 最大频率栈
实现 FreqStack，模拟类似栈的数据结构的操作的一个类。
FreqStack 有两个函数：
push(int x)，将整数 x 推入栈中。
pop()，它移除并返回栈中出现最频繁的元素。
如果最频繁的元素不只一个，则移除并返回最接近栈顶的元素。

示例：
输入：
["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
输出：[null,null,null,null,null,null,null,5,7,5,4]
解释：
执行六次 .push 操作后，栈自底向上为 [5,7,5,7,4,5]。然后：

pop() -> 返回 5，因为 5 是出现频率最高的。
栈变成 [5,7,5,7,4]。

pop() -> 返回 7，因为 5 和 7 都是频率最高的，但 7 最接近栈顶。
栈变成 [5,7,5,4]。

pop() -> 返回 5 。
栈变成 [5,7,4]。

pop() -> 返回 4 。
栈变成 [5,7]。
*/

type FreqStack struct {
	max    int
	numMap map[int]int
	cntMap map[int][]int
}

func Constructor() FreqStack {
	return FreqStack{
		numMap: make(map[int]int, 0),
		cntMap: make(map[int][]int, 0),
	}
}

func (this *FreqStack) Push(val int) {
	this.numMap[val]++
	if this.numMap[val] > this.max {
		this.max = this.numMap[val]
	}
	this.cntMap[this.numMap[val]] = append(this.cntMap[this.numMap[val]], val)
}

func (this *FreqStack) Pop() int {
	for this.max > 0 {
		list, ok := this.cntMap[this.max]
		if !ok || len(list) == 0 {
			this.max--
			continue
		}
		num := list[len(list)-1]
		this.cntMap[this.max] = list[:len(list)-1]
		this.numMap[num]--
		return num
	}
	return 0
}

func TestName(t *testing.T) {
	obj := Constructor()
	obj.Push(5)
	obj.Push(7)
	obj.Push(5)
	obj.Push(7)
	obj.Push(4)
	obj.Push(5)
	fmt.Println(obj.Pop()) // 5
	fmt.Println(obj.Pop()) // 7
	fmt.Println(obj.Pop()) // 5
	fmt.Println(obj.Pop()) // 4

	obj2 := Constructor()
	obj2.Push(1)
	obj2.Push(1)
	obj2.Push(1)
	obj2.Push(2)
	fmt.Println(obj2.Pop()) // 1
	fmt.Println(obj2.Pop()) // 1
	obj2.Push(2)
	obj2.Push(2)
	obj2.Push(1)
	fmt.Println(obj2.Pop()) // 2
	fmt.Println(obj2.Pop()) // 1
	fmt.Println(obj2.Pop()) // 2
}
