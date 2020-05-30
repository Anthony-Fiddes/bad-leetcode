package main

func combinationSum(candidates []int, target int) [][]int {

	var c = make([]int, target+1)
	coins := candidates
	var cc = make([][]int, len(coins))

	// Work on understanding why wayR works

	for i := range cc {
		cc[i] = make([]int, target+1)
	}

	for i := range cc {
		for j := range cc[i] {
			// For human reading purposes, totally unneccesary if statement
			if j == 0 {
				cc[i][j] = 1
				continue
			}
			cc[i][j] = -1
		}
	}

	// Finds permutations instead of combinations
	// like I need
	waysP := func(sum int) int {
		if sum < 0 {
			return 0
		} else if sum == 0 {
			return 1
		}

		if len(c) > sum && c[sum] != 0 {
			return c[sum]
		}

		var total int
		for _, v := range coins {
			total += waysP(sum - v)
		}
		c[sum] = total
		return total
	}

	Ways := func(sum int) int {
		return waysR(sum, 1)
	}

	index := func(coin int) int {
		for i := range coins {
			if coins[i] == coin {
				return i
			}
		}
		return -1
	}

	waysR := func(sum, coin int) int {
		// fmt.Println(sum, coin)
		coinIndex := index(coin)
		if cc[coinIndex][sum] >= 0 {
			return cc[coinIndex][sum]
		}

		if sum < 0 {
			return 0
		} else if sum == 0 {
			return 1
		}

		var total int
		for i, v := range coins {
			if v < coin {
				continue
			}

			coefficient := 1
			for coefficient*v <= sum {
				if i < len(coins)-1 {
					total += waysR(sum-coefficient*v, coins[i+1])
				} else {
					total += waysR(sum-coefficient*v, coins[len(coins)-1])
				}
				coefficient++
			}
		}
		cc[coinIndex][sum] = total
		return cc[coinIndex][sum]
	}

	return Ways(target)
}
