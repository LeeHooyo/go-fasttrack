# Day02 - Go Concurrency(goroutines)

## 1. goroutine
- `go` 키워드로 함수 실행하면 새로운 lightweight thread
- scheduler는 OS thread보다 훨씬 가벼움
- 순서는 보장되지 않음 (race 당연함)

## 2. sync.WaitGroup
- goroutine이 몇 개 생성되었는지 추적
- `wg.Add(1)`  → goroutine 시작 전에 추가
- `wg.Done()` → goroutine 종료 시 호출
- `wg.Wait()` → 모든 Done 완료될 때까지 block

## 3. 기본 패턴
```go
var wg sync.WaitGroup

for i := 0; i < n; i++ {
    wg.Add(1)
    go worker(&wg)
}

wg.Wait()
```

## 4. Rust와 대비 관점
- Rust: borrow checker + async/await 모델
- Go: GC 기반 + goroutine/channel 기본
- Rust는 “안전성의 강제”
- Go는 “빠른 생산성 + 단순함”

## 5. 기억할 것
- goroutine은 반드시 WaitGroup, channel 중 하나와 함께 써야 의미 있음
- time.Sleep은 demo에서만 사용 (실전에서는 거의 안씀)

-- 추가 --
## Channel 핵심

- 채널은 goroutine 간 데이터 전달 용도
- 타입이 지정된 파이프: `chan string`, `chan int`
- 방향 지정 가능:
  - `chan<- T` : send-only
  - `<-chan T` : recv-only
- close는 sender가 한다
- range는 close될 때까지 읽는다

## Channel 패턴 (기본)

```go
out := make(chan string, n)

go func() {
    out <- "hi"
}()
msg := <-out
```

## WaitGroup + Channel 조합의 의도

- WaitGroup = goroutine 종료 시점을 정확하게 조절
- Channel = 그 goroutine들의 결과를 안전하게 수집

Rust 대비:
- Rust는 borrow/ownership으로 동시성 안전을 강제
- Go는 동시성 자체는 풀려있고, Channel/WaitGroup으로 동기화
