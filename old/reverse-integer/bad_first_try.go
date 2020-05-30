package main

// import (
// 	"fmt"
// )

// func reverse(x int) int {
// 	var neg bool
// 	if x < 0 {
// 		x *= -1
// 		neg = true
// 	}
// 	d := make([]int32, 0)
// 	for x > 0 {
// 		d = append(d, int32(x%10))
// 		x /= 10
// 	}
// 	var r int32
// 	for i, v := range d {
// 		var last int32
// 		for j := len(d) - 1; j > i; j-- {
// 			last = v
// 			v = v * 10
// 			if v < last {
// 				return 0
// 			}
// 		}
// 		r += v
// 	}
// 	if r < 0 {
// 		return 0
// 	}
// 	if neg {
// 		r *= -1
// 	}
// 	return int(r)
// }

// func main() {
// 	fmt.Println(reverse(1563847412))
// }
