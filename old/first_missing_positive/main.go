package main

import "fmt"

func correctPosition(nums []int, i int) {
	val := nums[i]
	if val <= 0 || val > len(nums) || i == val-1 || val == nums[val-1] {
		return
	}
	nums[val-1], nums[i] = nums[i], nums[val-1]
	correctPosition(nums, i)
}

func firstMissingPositive(nums []int) int {
	expected := 1
	for i := range nums {
		correctPosition(nums, i)
	}
	for _, v := range nums {
		if expected != v {
			return expected
		}
		expected++
	}
	return expected
}

func main() {
	fmt.Println(firstMissingPositive([]int{}))
}
