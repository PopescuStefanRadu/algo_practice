package coin_change

import "sort"

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	return coinChangeRec(coins, amount, make(map[int]int, amount+1))
}

func coinChangeRec(coins []int, amount int, mem map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if count, ok := mem[amount]; ok {
		return count
	}

	maxValidPos := sort.SearchInts(coins, amount)
	if maxValidPos == len(coins) {
		maxValidPos--
	}
	if coins[maxValidPos] == amount {
		return 1
	}

	var min *int
	for i := maxValidPos; i >= 0; i-- {
		coinVal := coins[i]
		newAmount := amount - coinVal
		result := coinChangeRec(coins, newAmount, mem)
		mem[newAmount] = result

		if result < 0 {
			continue
		}

		if min == nil {
			min = &result
		} else {
			*min = minInt(result, *min)
		}
	}

	if min == nil {
		return -1
	}

	return *min + 1
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
