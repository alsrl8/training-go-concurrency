# 04. Practical Concurrency Patterns

## 목표

실전에서 자주 쓰이는 동시성 패턴을 직접 구현한다.

## 과제

`patterns.go`의 함수 4개를 구현하세요.

### 1. `WorkerPool(jobs []int, numWorkers int) []int`
- jobs 슬라이스의 각 값을 **제곱**하여 결과를 반환
- numWorkers개의 goroutine이 job channel에서 작업을 꺼내 처리
- 결과 순서는 무관
- 모든 작업 완료 후 정상 종료되어야 함 (goroutine 누수 없음)

### 2. `Pipeline(nums []int) []int`
- 3단계 파이프라인: **generate → double → add10**
- generate: 슬라이스의 값을 channel로 보냄
- double: 받은 값을 2배로 만들어 다음 channel로 보냄
- add10: 받은 값에 10을 더해 다음 channel로 보냄
- 최종 결과를 수집하여 반환 (순서 보장)

### 3. `RateLimitedFetch(urls []string, rps int) []FetchResult`
- 주어진 URL 목록에 HTTP GET 요청을 보내되, **초당 rps회**를 넘지 않도록 제한
- `time.Ticker`를 사용하여 rate limiting 구현
- 각 요청의 결과(URL, StatusCode, Error)를 수집하여 반환
- 순서는 입력과 동일하게 유지할 것

### 4. `FetchAll(ctx context.Context, urls []string) ([]FetchResult, error)`
- `golang.org/x/sync/errgroup`을 사용하여 URL 목록에 병렬 GET 요청
- 하나라도 실패하면 나머지를 **취소**하고 첫 번째 에러를 반환
- 모두 성공하면 결과를 반환 (순서는 입력과 동일)

## 실행

```bash
cd 04-patterns
go mod tidy
go test -v -race
```
