// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
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
	// Shift all elements to their correct position
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
