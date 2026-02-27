package sync

import "sync"

// SafeCounter는 여러 goroutine에서 안전하게 사용할 수 있는 카운터이다.
// sync.Mutex를 사용하여 동시 접근을 보호한다.
type SafeCounter struct {
	mutex sync.Mutex
	value int
}

// Increment는 카운터를 1 증가시킨다.
func (c *SafeCounter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

// Value는 현재 카운터 값을 반환한다.
func (c *SafeCounter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

// FetchOnce는 여러 goroutine이 동시에 Get()을 호출해도
// 실제 로딩 함수는 딱 한 번만 실행되는 구조체이다.
type FetchOnce struct {
	loader func() string
	once   sync.Once
	result string
}

// NewFetchOnce는 로딩 함수를 받아 FetchOnce를 생성한다.
func NewFetchOnce(loader func() string) *FetchOnce {
	return &FetchOnce{loader: loader}
}

// Get은 로딩된 데이터를 반환한다. 최초 호출 시에만 loader가 실행된다.
func (f *FetchOnce) Get() string {
	f.once.Do(func() {
		f.result = f.loader()
	})
	return f.result
}

// ParallelSum은 nums를 workers개로 분할하여 병렬로 합산한다.
// sync.WaitGroup으로 완료를 대기하고, 부분합을 안전하게 합산한다.
func ParallelSum(nums []int, workers int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) < workers {
		workers = len(nums)
	}

	chunkSize := len(nums) / workers
	startIndex, lastIndex := 0, 0 // range: [startIndex, lastIndex)
	wg := sync.WaitGroup{}
	ch := make(chan int, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)

		lastIndex += chunkSize
		if lastIndex+chunkSize > len(nums) {
			lastIndex = len(nums)
		}
		go func(start, last int) {
			defer wg.Done()
			tmp := 0
			for j := start; j < last; j++ {
				tmp += nums[j]
			}
			ch <- tmp
		}(startIndex, lastIndex)
		startIndex = lastIndex
	}

	wg.Wait()
	close(ch)

	var result int
	for v := range ch {
		result += v
	}
	return result
}
