package task1

func SingleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func SingleNumber1(nums []int) int {
	var ret int
	for _, v := range nums {
		ret ^= v
	}
	return ret
}
