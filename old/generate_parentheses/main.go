package main

func generateParenthesis(n int) []string {
	r := make([]string, 0)
	addParentheses(&r, "", n, 0, 0)
	return r
}

func addParentheses(combinations *[]string, curr string, n int, open int, close int) {
	if len(curr) == n*2 {
		*combinations = append(*combinations, curr)
		return
	}
	if open < n {
		addParentheses(combinations, curr+"(", n, open+1, close)
	}
	if close < open {
		addParentheses(combinations, curr+")", n, open, close+1)
	}
}

func main() {

}
