package enums

import (
	"fmt"
	"strconv"
)

func SplitDigits(num int64) []int {
	numStr := strconv.FormatInt(num, 10)
	fmt.Println(numStr)

	digits := make([]int, len(numStr))
	for i, digitRune := range numStr {
		digitInt, _ := strconv.Atoi(string(digitRune))
		digits[i] = digitInt
	}

	return digits
}

var PowersOfTwo = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

// returns the actual number and its index
func GetPowerOfTwoJustBelowNum(permLevel int) (int, int) {
	if permLevel <= 0 {
		return 0, 0
	}

	for i, power := range PowersOfTwo {
		if power > permLevel {
			return PowersOfTwo[i-1], i - 1
		}
		i++
	}
	return 1, 0
}
