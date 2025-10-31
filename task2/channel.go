package task2

import (
	"fmt"
	"time"
)

func Channel() {
	fmt.Println("task2 channel")
	ch := make(chan int)
	go sender(ch)
	go receiver(ch)
	time.Sleep(2 * time.Second)

	fmt.Println("task2 channel buffer")
	ch1 := make(chan int, 5)
	go bufferSender(ch1)
	go bufferReceiver(ch1)
	time.Sleep(2 * time.Second)

}

func sender(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func receiver(ch <-chan int) {
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println("")
}

func bufferSender(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func bufferReceiver(ch <-chan int) {
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println("")
}
