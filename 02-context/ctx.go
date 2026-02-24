package ctx

import (
	"context"
	"time"
)

// fetchWithTimeout은 주어진 URL에 timeout이 적용된 HTTP GET 요청을 보낸다.
// 성공 시 HTTP 상태 코드, 실패 시 에러를 반환한다.
func fetchWithTimeout(url string, timeout time.Duration) (int, error) {
	// TODO: 구현하세요
	_ = url
	_ = timeout
	return 0, nil
}

// doWork는 ctx가 취소될 때까지 1부터 시작하는 정수를 100ms 간격으로 results에 보낸다.
// ctx가 취소되면 results channel을 닫고 종료한다.
func doWork(ctx context.Context, results chan<- int) {
	// TODO: 구현하세요
}

// fanOutWithCancel은 n개의 goroutine을 실행하여 각각 자신의 인덱스를 반환한다.
// ctx가 취소되면 미완료 goroutine은 결과를 보내지 않는다.
// goroutine 누수가 없어야 한다.
func fanOutWithCancel(ctx context.Context, n int) []int {
	// TODO: 구현하세요
	return nil
}
