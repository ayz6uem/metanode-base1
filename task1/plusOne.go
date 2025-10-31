package task1

func PlusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	if len(digits) == 1 {
		if digits[0] == 9 {
			return []int{1, 0}
		}
	}
	if digits[0] == 9 {
		digits[0] = 0
		digits[len(digits)-1] += 1
		return digits
	}
	digits[len(digits)-1] += 1
	return digits
}

func PlusOne1(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	if len(digits) == 1 && digits[0] == 9 {
		return []int{1, 0}
	}
	f := false
	if digits[0] == 9 {
		f = true
		digits[len(digits)-1] = 0
	} else {
		digits[len(digits)-1] += 1
	}
	if f {
		digits[len(digits)-2] += 1
	}
	return digits
}
