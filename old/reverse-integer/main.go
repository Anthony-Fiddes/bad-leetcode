package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var neg bool
	if x < 0 {
		x *= -1
		neg = true
	}
	var r int32
	for x > 0 {
		if r > math.MaxInt32/10 {
			return 0
		}
		r = r*10 + int32(x%10)
		x /= 10
	}
	if neg {
		r *= -1
	}
	return int(r)
}

func main() {
	fmt.Println(reverse(1563847412))
}
