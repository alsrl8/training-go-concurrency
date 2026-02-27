package sync

// SafeCounter는 여러 goroutine에서 안전하게 사용할 수 있는 카운터이다.
// sync.Mutex를 사용하여 동시 접근을 보호한다.
type SafeCounter struct {
	// TODO: 필드를 정의하세요
}

// Increment는 카운터를 1 증가시킨다.
func (c *SafeCounter) Increment() {
	// TODO: 구현하세요
}

// Value는 현재 카운터 값을 반환한다.
func (c *SafeCounter) Value() int {
	// TODO: 구현하세요
	return 0
}

// FetchOnce는 여러 goroutine이 동시에 Get()을 호출해도
// 실제 로딩 함수는 딱 한 번만 실행되는 구조체이다.
type FetchOnce struct {
	// TODO: 필드를 정의하세요
}

// NewFetchOnce는 로딩 함수를 받아 FetchOnce를 생성한다.
func NewFetchOnce(loader func() string) *FetchOnce {
	// TODO: 구현하세요
	return nil
}

// Get은 로딩된 데이터를 반환한다. 최초 호출 시에만 loader가 실행된다.
func (f *FetchOnce) Get() string {
	// TODO: 구현하세요
	return ""
}

// ParallelSum은 nums를 workers개로 분할하여 병렬로 합산한다.
// sync.WaitGroup으로 완료를 대기하고, 부분합을 안전하게 합산한다.
func ParallelSum(nums []int, workers int) int {
	// TODO: 구현하세요
	return 0
}
