package ctx

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// fetchWithTimeout은 주어진 URL에 timeout이 적용된 HTTP GET 요청을 보낸다.
// 성공 시 HTTP 상태 코드, 실패 시 에러를 반환한다.
func fetchWithTimeout(url string, timeout time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

// doWork는 ctx가 취소될 때까지 1부터 시작하는 정수를 100ms 간격으로 results에 보낸다.
// ctx가 취소되면 results channel을 닫고 종료한다.
func doWork(ctx context.Context, results chan<- int) {
	defer close(results)

	num := 1
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(100 * time.Millisecond):
			results <- num
			num++
		}
	}
}

// fanOutWithCancel은 n개의 goroutine을 실행하여 각각 자신의 인덱스를 반환한다.
// ctx가 취소되면 미완료 goroutine은 결과를 보내지 않는다.
// goroutine 누수가 없어야 한다.
func fanOutWithCancel(ctx context.Context, n int) []int {
	ch := make(chan int, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			case ch <- idx:

			}
		}(i)
	}

	go func() { // 왜 이건 defer로 하지 않았지
		wg.Wait()
		close(ch)
	}()

	var results []int
	for v := range ch {
		results = append(results, v)
	}
	return results
}
