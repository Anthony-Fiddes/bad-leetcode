package main

import "fmt"

func pow10(exponent int) int {
	r := 1
	for i := 0; i < exponent; i++ {
		r *= 10
	}
	return r
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := x
	var p int
	for num != 0 {
		num /= 10
		p++
	}
	target := p / 2
	p--
	num = x
	for p >= target {
		pow := pow10(p)
		if x%10 != num/pow {
			return false
		}
		x /= 10
		num %= pow
		p--
	}
	return true
}

func main() {
	fmt.Println("0:", isPalindrome(0))
	fmt.Println("121:", isPalindrome(121))
	fmt.Println("123:", isPalindrome(123))
}
