package sync

import (
	gosync "sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	var wg gosync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()

	if got := c.Value(); got != 1000 {
		t.Errorf("counter = %d, want 1000", got)
	}
}

func TestFetchOnce(t *testing.T) {
	callCount := 0
	loader := func() string {
		callCount++
		return "loaded data"
	}

	f := NewFetchOnce(loader)
	var wg gosync.WaitGroup

	results := make([]string, 100)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			results[idx] = f.Get()
		}(i)
	}

	wg.Wait()

	if callCount != 1 {
		t.Errorf("loader called %d times, want 1", callCount)
	}

	for i, v := range results {
		if v != "loaded data" {
			t.Errorf("results[%d] = %q, want %q", i, v, "loaded data")
		}
	}
}

func TestParallelSum(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		workers int
		want    int
	}{
		{"empty", []int{}, 4, 0},
		{"single worker", []int{1, 2, 3, 4, 5}, 1, 15},
		{"two workers", []int{1, 2, 3, 4, 5}, 2, 15},
		{"more workers than items", []int{1, 2, 3}, 10, 6},
		{"large", makeRange(1, 1001), 8, 500500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParallelSum(tt.nums, tt.workers)
			if got != tt.want {
				t.Errorf("ParallelSum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func makeRange(min, max int) []int {
	nums := make([]int, max-min)
	for i := range nums {
		nums[i] = min + i
	}
	return nums
}
