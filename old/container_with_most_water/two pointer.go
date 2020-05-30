package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func area(height []int, i, j int) int {
	h := min(height[i], height[j])
	w := i - j
	if w < 0 {
		w *= -1
	}
	return h * w
}

func maxArea(height []int) int {
	a := 0
	b := len(height) - 1
	var max int
	for a < b {
		m := area(height, a, b)
		if m > max {
			max = m
		}
		if height[a] < height[b] {
			a++
		} else {
			b--
		}
	}

	return max
}

func main() {
	// maxArea([]int{1, 3, 2, 5, 25, 24, 5})
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
