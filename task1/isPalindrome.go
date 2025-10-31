package task1

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func IsPalindrome1(x int) bool {
	if x <= 0 || x%10 == 0 {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}
	return x == y || x == y/10
}
