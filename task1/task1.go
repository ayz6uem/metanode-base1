package task1

import "fmt"

func Task1() {
	fmt.Println("task1 只出现一次的数字:", SingleNumber([]int{2, 2, 1}))
	fmt.Println("task1 只出现一次的数字:", SingleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println("task1 只出现一次的数字:", SingleNumber([]int{1}))

	fmt.Println("task1 回文数:", IsPalindrome(121))
	fmt.Println("task1 回文数:", IsPalindrome(-121))
	fmt.Println("task1 回文数:", IsPalindrome(10))

	fmt.Println("task1 字符串有效的括号:", ValidBracket("()"))
	fmt.Println("task1 字符串有效的括号:", ValidBracket("()[]{}"))
	fmt.Println("task1 字符串有效的括号:", ValidBracket("(]"))
	fmt.Println("task1 字符串有效的括号:", ValidBracket("([])"))
	fmt.Println("task1 字符串有效的括号:", ValidBracket("([)]"))

	fmt.Println("task1 最长公共前缀:", LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("task1 最长公共前缀:", LongestCommonPrefix([]string{"dog", "racecar", "car"}))

	fmt.Println("task1 加一:", PlusOne([]int{1, 2, 3}))
	fmt.Println("task1 加一:", PlusOne([]int{4, 3, 2, 1}))
	fmt.Println("task1 加一:", PlusOne([]int{9}))

	fmt.Println("task1 删除有序数组中的重复项:", RemoveDuplicates([]int{1, 1, 2}))
	fmt.Println("task1 删除有序数组中的重复项:", RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	fmt.Println("task1 合并区间:", Merge1([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println("task1 合并区间:", Merge1([][]int{{1, 10}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println("task1 合并区间:", Merge1([][]int{{1, 4}, {4, 5}}))
	fmt.Println("task1 合并区间:", Merge1([][]int{{4, 7}, {1, 4}}))

	fmt.Println("task1 两数之和:", TwoSum1([]int{2, 7, 11, 15}, 9))
	fmt.Println("task1 两数之和:", TwoSum1([]int{3, 2, 4}, 6))
	fmt.Println("task1 两数之和:", TwoSum1([]int{3, 3}, 6))
}
