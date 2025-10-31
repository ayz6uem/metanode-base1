package task2

import (
	"fmt"
	"time"
)

func Goroutine() {
	goNumber()
	goSchedule()
}

func goNumber() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println("task2 协程 奇数", i)
			}
		}
	}()
	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("task2 协程 偶数", i)
			}
		}
	}()
	time.Sleep(time.Second)
}

func goSchedule() {
	schedule([]func(){
		func() {
			fmt.Println("task2 协程 开始执行方法")
			time.Sleep(2 * time.Second)
			fmt.Println("task2 协程 结束执行方法")
		},
		func() {
			fmt.Println("task2 协程 仅执行方法")
		},
	})
}

func schedule(fs []func()) {
	for _, f := range fs {
		go Runnable(f)
	}
	time.Sleep(3 * time.Second)
}

func Runnable(f func()) {
	startTime := time.Now()
	f()
	fmt.Println("task2 协程 执行时间:", time.Now().Sub(startTime))
}
