package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	var (
		min = -(1 << 31)
		max = (1 << 31) - 1
		y   int
	)

	//超过32位, 返回0
	if x <= min || x >= max || x == 0 {
		return 0
	}

	for x != 0 {
		//y*10 目的是前进一位, x%10 目的是取出最后一位
		//y*10 + x%10, 目的是拼接两个数字
		y = y*10 + x%10

		//超过32位, 返回0
		if y <= min || y >= max {
			return 0
		}

		//去除最后一位
		x /= 10
	}

	return y
}
