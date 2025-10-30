# Day01 - Go 기본 문법 요약

## 1. 핵심 문법 개요

| 개념 | 설명 | 예시 |
|------|------|------|
| 프로그램 진입점 | `package main` + `func main()` | `func main() { fmt.Println("Hi") }` |
| 변수 선언 | `var x int = 10` 또는 `x := 10` | `name := "yunkyu"` |
| 상수 | `const Pi = 3.14` | - |
| 타입 추론 | 오른쪽 값으로 자동 결정 | `x := 42` |
| 포인터 | `&x` 주소, `*p` 역참조 | `p := &x; fmt.Println(*p)` |
| 함수 | `func add(a int, b int) int { return a + b }` | - |
| 여러 리턴값 | Go의 대표적 특징 | `func div(a,b int)(int,error)` |
| defer | 함수 종료 직전 실행 예약 | `defer f.Close()` |
| if / else | 괄호 없음 | `if x > 10 { ... } else { ... }` |
| for | 유일한 반복문 | `for i := 0; i < n; i++ {}` |
| range | 슬라이스/맵 순회 | `for i,v := range arr {}` |
| switch | Rust의 match보다 단순 | `switch x { case 1: ... }` |
| array | `[5]int` 고정 길이 | - |
| slice | `[]int` 가변 길이 | `nums := []int{1,2,3}` |
| map | 키-값 쌍 | `m := map[string]int{"a":1}` |
| struct | 커스텀 타입 | `type User struct { Name string; Age int }` |
| interface | 동적 타입 규약 | `type Writer interface { Write([]byte)(int,error) }` |
| error | 내장 인터페이스 | `if err != nil { ... }` |

---

## 2. 연산자와 비교식

| 분류 | 기호 | 예시 |
|------|------|------|
| 대입 | `=` | `x = 3` |
| 복합 대입 | `+= -= *= /= %=` | - |
| 선언 + 대입 | `:=` | `x := 10` |
| 비교 | `== != < > <= >=` | - |
| 논리 | `&& || !` | - |
| 증감 | `i++ i--` | 독립문으로만 사용 가능 |
| 주소/역참조 | `&x`, `*p` | - |

---

## 3. 출력 포맷 (fmt 패키지)

| 함수 | 설명 | 예시 |
|------|------|------|
| `fmt.Println` | 자동 개행 | `fmt.Println("hi", x)` |
| `fmt.Printf` | 포맷 지정 | `fmt.Printf("num=%d str=%s\n", n, s)` |
| `fmt.Sprintf` | 문자열 반환 | `msg := fmt.Sprintf("x=%d", x)` |

| 포맷 | 의미 |
|------|------|
| `%d` | 정수 |
| `%f` | 실수 |
| `%s` | 문자열 |
| `%v` | 자동 포맷 (debug용) |
| `%T` | 타입 표시 |

---

## 4. 슬라이스 / 배열 / 맵 기본

```
nums := []int{1,2,3}        // 슬라이스
nums = append(nums, 4, 5)   // 요소 추가

arr := [3]int{1,2,3}        // 고정 배열
fmt.Println(len(arr))       // 길이

m := map[string]int{"a":1,"b":2}
m["c"] = 3
delete(m, "b")
for k, v := range m { fmt.Println(k, v) }
```

---

## 5. 함수와 에러 처리

```
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

res, err := divide(10, 2)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("Result:", res)
}
```

---

## 6. defer, panic, recover

```
func main() {
	defer fmt.Println("끝!") // 함수 종료 시 실행
	fmt.Println("시작")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("복구:", r)
		}
	}()

	panic("치명적 오류 발생!") // 강제 중단 → recover로 복구 가능
}
```

---

## 7. 구조체와 메서드

```
type Account struct {
	Name    string
	Balance int
}

// 메서드 정의
func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func main() {
	acc := Account{Name: "Lee", Balance: 1000}
	acc.Deposit(500)
	fmt.Printf("%+v\n", acc)
}
```

---

## 8. import와 package 구조

- 모든 Go 파일은 `package`로 시작한다.  
- 동일 폴더 내 파일은 같은 package를 공유한다.  
- 다른 패키지를 가져올 때는 모듈 경로 사용:
  ```go
  import "go-fasttrack/labs/utils"
  ```
- 외부 패키지:
  ```bash
  go get github.com/gin-gonic/gin
  ```

---

## 9. Go의 철학 요약

- 변수는 선언 후 반드시 사용해야 한다.  
- 들여쓰기는 탭 1개 (`go fmt`로 강제됨).  
- 예외(Exception) 없음 — 에러는 값으로 다룸.  
- 가비지 컬렉션 있음 (Rust처럼 소유권 시스템 없음).  
- 간결함을 최우선으로 하는 설계.

---

## 10. 이걸로 커버 가능한 영역

- CLI 프로그램 (`fmt`, `os`, `flag`)
- 파일 IO (`os.Open`, `os.WriteFile`)
- HTTP 서버 (`net/http`)
- JSON 직렬화 (`encoding/json`)
- Goroutine / channel (`go func()`, `make(chan int)`)
- 기본 자료구조 / 알고리즘 구현

---

## 11. 핵심 정리 템플릿

```md
# 핵심 문법
- 변수 선언: `x := 10`, `var x int = 10`
- 함수: `func add(a,b int) int`
- 포인터: `p := &x`, `*p`
- 슬라이스: `nums := []int{1,2,3}`
- 맵: `m := map[string]int{"a":1}`
- 출력: `fmt.Println`, `fmt.Printf("%d", x)`
- 조건문: `if`, `switch`
- 반복문: `for`, `range`
- 에러 처리: `if err != nil`
- defer / panic / recover
- 구조체와 메서드: `type User struct {...}` + `(u *User) Greet()`

# 기억할 점
- 변수 미사용 금지
- 함수는 명시적 리턴
- 에러는 값으로 다룸 (예외 아님)
- `go fmt`로 스타일 강제
```

