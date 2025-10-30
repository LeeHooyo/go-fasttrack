
# Go Quick Reference

## Basics
- build & run  
  ~~~
  go run .
  go build
  ~~~
- format & test  
  ~~~
  go fmt ./...
  go test ./...
  ~~~
- tidy up modules  
  ~~~
  go mod tidy
  ~~~

## Packages
- `"fmt"`, `"os"`, `"io"`, `"bufio"`
- `"encoding/json"`, `"net/http"`
- `"context"`, `"errors"`

## Concurrency
- goroutine: `go func()`
- channel: `make(chan T)`
- select: multiplex goroutines
- context: cancellation & timeout

## Tips
- 모든 Go 프로그램은 `package main` + `func main()` 필요
- 모듈 초기화: `go mod init <module>`
- 워크스페이스 연결: `go work init ./labs/...`
- 실행: `go run ./labs/...`

## Common Errors
| 에러 | 원인 | 해결 |
|------|------|------|
| no Go files | main.go 없음 | main.go 추가 |
| no modules found | go.work 누락 | go work init 다시 |
| undefined symbol | import 누락 | go fmt ./... or go mod tidy |

