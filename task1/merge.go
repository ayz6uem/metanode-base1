package task1

func Merge(intervals [][]int) [][]int {
	f := false
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if isCanMerge(intervals[i], intervals[j]) {
				f = true
				intervals[i] = []int{min(intervals[i][0], intervals[j][0]), max(intervals[i][1], intervals[j][1])}
				intervals = append(intervals[:j], intervals[j+1:]...)
				break
			}
		}
	}
	if f {
		return Merge(intervals)
	}
	return intervals
}

func Merge1(intervals [][]int) [][]int {
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	result := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		f := false
		for j := len(intervals) - 1; j > i; j-- {
			if isCanMerge(intervals[i], intervals[j]) {
				f = true
				result = append(result, []int{intervals[i][0], intervals[j][1]})
				i = j
				break
			}
		}
		if !f {
			result = append(result, intervals[i])
		}
	}
	return result
}

func isCanMerge(arr1 []int, arr2 []int) bool {
	return (arr1[0] <= arr2[1] && arr1[1] >= arr2[0]) || (arr2[0] <= arr1[1] && arr2[1] >= arr1[0])
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
