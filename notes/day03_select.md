# Day03 - select를 통한 채널 멀티플렉싱

## 1. select란?
- 여러 채널의 상태를 동시에 감시하고, 준비된 케이스를 즉시 실행.
- non-blocking 비동기 처리를 단순하게 만든다.
- 모든 채널이 block된 경우 select 자체가 block됨.

## 2. 기본 구조
```go
select {
case msg := <-ch1:
    fmt.Println(msg)
case msg := <-ch2:
    fmt.Println(msg)
default:
    fmt.Println("no message ready") // optional
}
```

## 3. 특징
- `default`를 넣으면 block 없이 즉시 넘어감.
- 닫힌 채널은 즉시 zero value를 반환.
- nil 채널은 select에서 완전히 제외됨 (활성화 제어에 사용).

## 4. time.After / time.Tick
- `time.After(d)` : d 후에 신호 1회
- `time.Tick(d)` : d마다 주기적 신호
```go
select {
case <-time.After(5 * time.Second):
    fmt.Println("timeout")
}
```

## 5. 실무 패턴
- timeout / cancel 처리
- 여러 worker 결과 중 가장 빠른 것 우선 사용
- streaming API / event listener

## 6. Rust와 비교
| 개념 | Rust | Go |
|------|------|----|
| 비동기 모델 | async/await + Future | goroutine + channel + select |
| 스케줄링 | 명시적 executor | 런타임 내장 |
| 안전성 | 컴파일 타임 보장 | 런타임 제어, 단순함 |
| 철학 | 안전성 우선 | 단순성 + 실용성 우선 |

## 7. 정리
- goroutine + channel + select = Go의 동시성 3요소
- select는 channel 기반 비동기 로직을 구성하는 기본 단위
- timeout, cancel, multiplex 패턴을 여기서 전부 구현 가능

