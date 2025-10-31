package task1

func LongestCommonPrefix(strs []string) string {
	var first = strs[0]
	var index = 0
outer:
	for ; index < 200; index++ {
		if len(first) <= index {
			break outer
		}
		r := first[index]
		for _, str := range strs {
			if len(str) <= index {
				break outer
			}
			if r != str[index] {
				break outer
			}
		}
	}

	var result []rune = []rune(first)
	return string(result[:index])
}
