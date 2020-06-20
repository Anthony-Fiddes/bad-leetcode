// https://leetcode.com/problems/remove-element/
package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	last := len(nums) - 1
	for i < last {
		if nums[i] == val {
			nums[last], nums[i] = nums[i], nums[last]
			last--
			continue
		}
		i++
	}
	for i, v := range nums {
		if v == val {
			nums = nums[:i]
			return len(nums)
		}
	}
	return len(nums)
}

func main() {
	removeElement([]int{1}, 1)
}
