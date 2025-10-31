package task1

import "fmt"

func RemoveDuplicates(nums []int) int {
	i, j := 1, 0
	for ; i < len(nums); i++ {
		if nums[i] != nums[j] {
			nums[j+1] = nums[i]
			j = j + 1
		}
	}
	fmt.Println(nums)
	return j + 1
}
