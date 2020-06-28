// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Runtime: 116 ms, faster than 5.35% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 4.6 MB, less than 99.88% of Go online submissions for Remove Duplicates from Sorted Array.
package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	min := nums[0]
	last := min
	// Set all duplicates to be the minimum element
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		if val == last {
			nums[i] = min
		} else {
			last = val
		}
	}
	// Shift all minimum elements to the end of the array
	for i := 1; i < len(nums); i++ {
		if nums[i] == min {
			next := i + 1
			for next < len(nums)-1 && nums[next] == min {
				next++
			}
			if next >= len(nums) {
				break
			}
			nums[i], nums[next] = nums[next], nums[i]
		}
	}
	// Determine the true length of the array
	var length int
	repeats := false
	for i := 1; i < len(nums); i++ {
		if nums[i] == min {
			length = i
			nums = nums[:length]
			repeats = true
			break
		}
	}
	if !repeats {
		length = len(nums)
	}
	return length
}

func main() {}
