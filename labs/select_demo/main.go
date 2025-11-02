package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 두 개의 작업이 비동기적으로 결과를 보내는 상황
func worker(name string, ch chan<- string) {
	d := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(d)
	ch <- fmt.Sprintf("%s finished after %v", name, d)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan string)
	ch2 := make(chan string)

	go worker("alpha", ch1)
	go worker("beta", ch2)

	timeout := time.After(600 * time.Millisecond)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("[recv]", msg1)
			ch1 = nil // 채널 nil 처리 → 이후 select에서 제외
		case msg2 := <-ch2:
			fmt.Println("[recv]", msg2)
			ch2 = nil
		case <-timeout:
			fmt.Println("[timeout] 작업이 너무 오래 걸림")
			return
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}

	fmt.Println("select: done")
}

