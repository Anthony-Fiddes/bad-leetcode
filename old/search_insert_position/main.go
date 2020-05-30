package main

import "fmt"

func binarySearch(nums []int, target, low, high int) (int, error) {
	m := (low + high) / 2
	if low >= high {
		err := fmt.Errorf("target not found")
		if target <= nums[m] {
			return m, err
		}
		return m + 1, fmt.Errorf("target not found")
	}
	if nums[m] < target {
		return binarySearch(nums, target, m+1, high)
	} else if nums[m] > target {
		return binarySearch(nums, target, low, m-1)
	}
	return m, nil
}

func searchInsert(nums []int, target int) int {
	t, _ := binarySearch(nums, target, 0, len(nums)-1)
	return t
}

func main() {}
