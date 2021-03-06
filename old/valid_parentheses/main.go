package main

import "fmt"

var opener = map[rune]rune{
	')': '(',
	'}': ':',
	']': '[',
}

type stack []rune

func (s *stack) Pop() (rune, error) {
	r, err := s.Peek()
	if err != nil {
		return 0, err
	}
	*s = (*s)[:len(*s)-1]
	return r, nil
}

func (s stack) Peek() (rune, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("stack empty")
	}
	return s[len(s)-1], nil
}

func (s *stack) Push(r rune) {
	*s = append(*s, r)
}

func isValid(s string) bool {
	b := make(stack, 0)
	for _, v := range s {
		if len(b) > 0 {
			top, err := b.Peek()
			if err != nil {
				fmt.Println(err)
				return false
			}
			if top == opener[v] {
				b.Pop()
				continue
			}
		}
		b.Push(v)
	}
	return len(b) == 0
}

func main() {}
