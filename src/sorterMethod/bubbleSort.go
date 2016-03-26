package sorterMethod

import (
// "fmt"
)

func BubbleSort(input []int) {
	if nil == input || len(input) == 0 {
		return
	}
	length := len(input)
	for i := 0; i < length-1; i++ {
		flag := true
		for j := 0; j < length-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
