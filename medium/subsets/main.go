// https://leetcode.com/problems/subsets/
package main

import "fmt"

func subsets(nums []int) [][]int {
	cap := 1
	for i := 0; i < len(nums); i++ {
		cap *= 2
	}
	r := make([][]int, 0, cap)
	r = append(r, []int{})
	addSubset(&r, nums, []int{})
	return r
}

func addSubset(subsets *[][]int, nums []int, curr []int) {
	if len(nums) == 0 {
		return
	}
	for i, v := range nums {
		curr = append(curr, v)
		subset := make([]int, len(curr))
		copy(subset, curr)
		*subsets = append(*subsets, subset)
		addSubset(subsets, nums[i+1:], curr)
		curr = curr[:len(curr)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}
