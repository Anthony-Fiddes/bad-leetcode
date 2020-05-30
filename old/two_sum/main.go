package main

import "fmt"

func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int, len(nums))
	for i, v := range nums {
		o, ok := diffs[v]
		if ok {
			return []int{o, i}
		}
		diffs[target-v] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
