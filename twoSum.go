package main

import "fmt"

func main() {
	nums := make([]int, 4)
	nums[0] = 0
	nums[1] = 4
	nums[2] = 3
	nums[3] = 0
	fmt.Println(twoSum(nums, 7))
	fmt.Println(twoSum2(nums, 7))
}

func twoSum(nums []int, target int) []int {
	for k1, v1 := range nums {

		num := len(nums) - k1 - 1
		for i := 1; i <= num; i++ {
			if (v1 + nums[k1+i]) == target {
				return []int{k1, k1 + 1}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}

	for i, v := range nums {
		//取出目标值和当前value的差值, 判断map中是否存在
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}

		//反转key和vale
		//此处代码只能放在最后面, 防止出现重复值
		m[v] = i
	}

	return nil
}
