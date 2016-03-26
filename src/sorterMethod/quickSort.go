package sorterMethod

import (
// "fmt"
)

func QuickSort(input []int) {
	if input == nil || len(input) == 0 {
		return
	}
	length := len(input)
	quickSortSub(input, 0, length-1)
}

func quickSortSub(input []int, left, right int) {
	i := left
	j := right
	p := left
	// temp := input[left]
	for i <= j {
		for p <= j && input[p] <= input[j] {
			j--
		}
		if p <= j {
			input[p], input[j] = input[j], input[p]
			p = j
		}
		for p >= i && input[p] >= input[i] {
			i++
		}
		if p >= i {
			input[p], input[i] = input[i], input[p]
			p = i
		}
	}
	if p-left > 1 {
		//非单元素集合,继续细分
		quickSortSub(input, left, p-1)
	}
	if right-p > 1 {
		//非单元素集合,继续细分
		quickSortSub(input, p+1, right)
	}
}
