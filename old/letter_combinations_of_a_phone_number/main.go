package main

var chars = make(map[byte]string, 8)

func init() {
	chars['2'] = "abc"
	chars['3'] = "def"
	chars['4'] = "ghi"
	chars['5'] = "jkl"
	chars['6'] = "mno"
	chars['7'] = "pqrs"
	chars['8'] = "tuv"
	chars['9'] = "wxyz"
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	r := []string{""}
	r = addCombos(r, digits)
	return r
}

func addCombos(combos []string, digits string) []string {
	if digits == "" {
		return combos
	}
	curr := chars[digits[0]]
	r := make([]string, 0, len(combos)*len(curr))
	for _, v := range combos {
		for _, c := range curr {
			r = append(r, v+string(c))
		}
	}
	return addCombos(r, digits[1:])
}

func main() {
	letterCombinations("23457")
}
