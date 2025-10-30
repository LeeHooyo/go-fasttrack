# go-fasttrack

> Goal: Learn Go fast and break into open-source / blockchain development within 3 months.

---

## Structure
~~~
~/Projects/go-fasttrack/
├── go.work
├── notes/
│   └── quick_ref.md
├── labs/
│   ├── basic_syntax/
│   ├── goroutines/
│   ├── http_server/
│   └── rest_api/
└── projects/
    └── mini_wallet/
~~~

---

## Quick Start
~~~
# Run labs (once main.go added)
go run ./labs/basic_syntax
go run ./labs/goroutines
go run ./labs/http_server        # http://localhost:8080
go run ./labs/rest_api           # http://localhost:8081/health

# Run project
go run ./projects/mini_wallet --amount 1000 --memo "init"
~~~

---

## Conventions
- Go 1.22+
- Code style: `go fmt ./...`
- No external dependencies (std only until Phase 2)
- Each lab = independent runnable module

---

## Milestones
| Phase | Period | Goal |
|-------|---------|------|
| 1 | Week 1–2 | Go fundamentals (syntax, concurrency, http) |
| 2 | Week 3–6 | Blockchain OSS analysis & contribution |
| 3 | Week 7–10 | Mini service project |
| 4 | Week 11–12 | Portfolio finalization & job applications |

---

## Notes
All language notes, snippets, and commands go into `notes/quick_ref.md`.
