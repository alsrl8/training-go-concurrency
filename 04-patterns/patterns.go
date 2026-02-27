package patterns

import (
	"context"
	"time"
)

// FetchResult는 HTTP 요청 결과를 담는 구조체이다.
type FetchResult struct {
	URL        string
	StatusCode int
	Err        error
}

// WorkerPool은 jobs의 각 값을 제곱하여 반환한다.
// numWorkers개의 goroutine이 job channel에서 작업을 꺼내 처리한다.
func WorkerPool(jobs []int, numWorkers int) []int {
	// TODO: 구현하세요
	_ = jobs
	_ = numWorkers
	return nil
}

// Pipeline은 3단계 파이프라인(generate → double → add10)으로 값을 변환한다.
// 입력 순서가 유지되어야 한다.
func Pipeline(nums []int) []int {
	// TODO: 구현하세요
	_ = nums
	return nil
}

// RateLimitedFetch는 초당 rps회를 넘지 않도록 HTTP GET 요청을 보낸다.
// time.Ticker를 사용하여 rate limiting을 구현한다.
// 결과 순서는 입력과 동일해야 한다.
func RateLimitedFetch(urls []string, rps int) []FetchResult {
	// TODO: 구현하세요
	_ = urls
	_ = rps
	_ = time.Now()
	return nil
}

// FetchAll은 errgroup을 사용하여 URL 목록에 병렬 GET 요청을 보낸다.
// 하나라도 실패하면 나머지를 취소하고 에러를 반환한다.
func FetchAll(ctx context.Context, urls []string) ([]FetchResult, error) {
	// TODO: 구현하세요
	_ = ctx
	_ = urls
	return nil, nil
}
