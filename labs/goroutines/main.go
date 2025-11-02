package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	out <- fmt.Sprintf("worker %d done", id)
}

func main() {
	const n = 5

	var wg sync.WaitGroup
	out := make(chan string, n) // 버퍼 채널, n개까지 블로킹 X

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, out, &wg)
	}

	// goroutines 모두 종료될 때까지 기다림
	wg.Wait()
	close(out) // 반드시 닫아야 range가 끝남

	// 채널에서 수집
	for msg := range out {
		fmt.Println(msg)
	}

	fmt.Println("all done")
}
