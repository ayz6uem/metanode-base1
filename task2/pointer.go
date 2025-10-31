package task2

import "fmt"

func Pointer() {
	i := 1
	pointerAdd10(&i)
	fmt.Println("task2 指针:", i)

	arr := []int{1, 2, 3}
	pointerSlice(&arr)
	fmt.Println("task2 指针:", arr)
}

func pointerAdd10(p *int) {
	*p += 10
}

func pointerSlice(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}
