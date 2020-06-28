// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Runtime: 8 ms, faster than 88.76% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 4.6 MB, less than 99.88% of Go online submissions for Remove Duplicates from Sorted Array.
package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i, j := 0, 1
	// j moves ahead of i to find the next unique number and put it at
	// the i+1 position. i is incremented after a unique number is found, so
	// this makes the ith index the location of the last unique number.
	for j < len(nums) {
		if nums[j] == nums[i] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
		}
	}
	nums = nums[:i+1]
	return i + 1
}

func main() {}
