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
	for j := i; j < len(nums); j++ {
		if nums[j] == val {
			nums = nums[0:j]
			return len(nums)
		}
	}
	return len(nums)
}

func main() {
	removeElement([]int{1}, 1)
}
