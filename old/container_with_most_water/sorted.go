package main

// import (
// 	"fmt"
// 	"sort"
// )

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// func maxi(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func area(height []int, i, j int) int {
// 	h := min(height[i], height[j])
// 	w := i - j
// 	if w < 0 {
// 		w *= -1
// 	}
// 	return h * w
// }

// func maxArea(height []int) int {
// 	var max int
// 	index := make(map[int][]int, len(height))
// 	for i, v := range height {
// 		index[v] = append(index[v], i)
// 	}
// 	s := make([]int, len(height))
// 	copy(s, height)
// 	sort.Ints(s)
// 	h := make([]int, len(height))
// 	copy(h, height)
// 	a := 0
// 	b := len(h) - 1
// 	for _, v := range s {
// 		for _, curr := range index[v] {
// 			m := maxi(area(height, curr, a), area(height, curr, b))
// 			if m > max {
// 				max = m
// 			}
// 			h[curr] = 0
// 			switch curr {
// 			case a:
// 				for a < len(h)-1 && h[a] == 0 {
// 					a++
// 				}
// 			case b:
// 				for b > 0 && h[b] == 0 {
// 					b--
// 				}
// 			}
// 		}
// 	}
// 	return max
// }

// func main() {
// 	// maxArea([]int{1, 3, 2, 5, 25, 24, 5})
// 	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// }
