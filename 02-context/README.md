# 02. Context

## 목표

`context` 패키지를 사용하여 goroutine의 취소(cancellation)와 타임아웃을 제어하는 방법을 익힌다.

## 과제

`ctx.go`의 함수 3개를 구현하세요.

### 1. `fetchWithTimeout(url string, timeout time.Duration) (int, error)`
- 주어진 URL에 HTTP GET 요청을 보내되, timeout 내에 응답이 없으면 에러를 반환
- `context.WithTimeout`을 사용하여 `http.NewRequestWithContext`로 요청을 생성할 것
- 성공 시 HTTP 상태 코드를 반환, 실패 시 에러를 반환
- response body는 반드시 닫을 것

### 2. `doWork(ctx context.Context, results chan<- int)`
- 1부터 시작하여 값을 results channel에 하나씩 보냄
- 각 값을 보내기 전에 100ms 대기
- **ctx가 취소되면 즉시 중단**하고 channel을 닫을 것
- `select`와 `time.After` 또는 `time.NewTicker`를 활용

### 3. `fanOutWithCancel(ctx context.Context, n int) []int`
- n개의 goroutine을 실행하여 각각 자신의 인덱스(0~n-1)를 반환
- 단, ctx가 취소되면 아직 완료되지 않은 goroutine은 결과를 보내지 않아야 함
- 취소 여부와 관계없이 goroutine 누수가 없어야 함 (`go test -race`로 검증)
- 수집된 결과만 반환 (순서 무관)

## 실행

```bash
cd 02-context
go test -v -race
```
