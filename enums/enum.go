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
