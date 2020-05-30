package main

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// func area(i, j, h int) int {
// 	w := i - j
// 	if w < 0 {
// 		w *= -1
// 	}
// 	return h * w
// }

// func maxArea(height []int) int {
// 	var max int
// 	for i := range height {
// 		var maxHeight int
// 		for j := len(height) - 1; j >= i+1; j-- {
// 			h := min(height[i], height[j])
// 			if h > maxHeight {
// 				maxHeight = h
// 			} else {
// 				continue
// 			}
// 			a := area(i, j, h)
// 			if a > max {
// 				max = a
// 			}
// 		}
// 	}
// 	return max
// }

// // func main() {}
