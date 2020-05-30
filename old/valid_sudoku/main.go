package main

func isValidRow(b [][]byte, r int) bool {
	return checker(b, r, -1, rowValidator)
}

func isValidCol(b [][]byte, c int) bool {
	return checker(b, -1, c, colValidator)
}

func isValidSquare(b [][]byte, x int, y int) bool {
	return checker(b, x-1, y, squareValidator(x, y))
}

// admittedly a naive, ugly approach, but it reduced code rewriting
// and was a learning experience. I saw a better solution that just
// used row and column bounds (e.g. check from (a,b) to (i,j))
func checker(b [][]byte, x, y int, f func(*int, *int) bool) bool {
	digits := make(map[byte]bool, 10)
	for f(&x, &y) {
		v := b[x][y]
		if v == '.' {
			continue
		}
		in := digits[v]
		if !in {
			digits[v] = true
		} else {
			return false
		}
	}
	return true
}

func rowValidator(x, y *int) bool {
	if *y < 8 {
		(*y)++
		return true
	}
	return false
}

func colValidator(x, y *int) bool {
	if *x < 8 {
		(*x)++
		return true
	}
	return false
}

func squareValidator(startX, startY int) func(x, y *int) bool {
	return func(x, y *int) bool {
		if *x < startX+2 {
			(*x)++
			return true
		}
		if *x == startX+2 && *y < startY+2 {
			*x -= 2
			(*y)++
			return true
		}
		return false
	}
}

// maybe try implementing concurrency
func isValidSudoku(board [][]byte) bool {
	for i := range board {
		if !isValidRow(board, i) {
			return false
		}
		if !isValidCol(board, i) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !isValidSquare(board, i, j) {
				return false
			}
		}
	}
	return true
}

// [
//   ["8","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]

func main() {
	// fmt.Println(isValidRow([]byte{8, 3, 3, '.', 7, '.', '.', '.', '.'}))
}
