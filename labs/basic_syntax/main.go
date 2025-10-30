package main

import (
	"fmt"
	"os"
)

// add: 두 수의 합을 반환
func add(a, b int) int {
	return a + b
}

// pointerDemo: 포인터를 사용한 값 변경 예시
func pointerDemo() {
	x := 10
	ptr := &x
	fmt.Println("[Before]", x)
	*ptr = 20
	fmt.Println("[After ]", x)
}

// sliceDemo: 슬라이스와 for-range 문법 예시
func sliceDemo() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for i, n := range nums {
		sum += n
		fmt.Printf("idx:%d val:%d\n", i, n)
	}
	fmt.Println("sum =", sum)
}

func main() {
	fmt.Println("=== Go Basic Syntax ===")

	// 1. 커맨드라인 인자 읽기
	if len(os.Args) > 1 {
		fmt.Println("Args:", os.Args[1:])
	} else {
		fmt.Println("No args provided.")
	}

	// 2. 함수 호출
	result := add(5,3)
	fmt.Println("add(5, 3) =", result)

	// 3. 포인터 데모
	pointerDemo()

	// 4. 슬라이스 데모
	sliceDemo()

	// 5. 에러 처리 기본
	f, err := os.Open("nonexistent.txt")
	if err != nil {
		fmt.Println("[Error handled]:", err)
	} else {
		defer f.Close()
	}
}
