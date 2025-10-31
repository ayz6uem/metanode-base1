package task1

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TwoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			m[nums[i]] = i
		} else {
			k, ok := m[target-nums[i]]
			if ok {
				return []int{k, i}
			} else {
				m[nums[i]] = i
			}
		}
	}
	return []int{}
}
