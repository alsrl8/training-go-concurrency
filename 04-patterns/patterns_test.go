package patterns

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := WorkerPool(jobs, 3)

	if len(results) != len(jobs) {
		t.Fatalf("got %d results, want %d", len(results), len(jobs))
	}

	sort.Ints(results)
	want := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	for i, v := range results {
		if v != want[i] {
			t.Errorf("results[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestWorkerPool_Empty(t *testing.T) {
	results := WorkerPool([]int{}, 3)
	if len(results) != 0 {
		t.Errorf("got %d results, want 0", len(results))
	}
}

func TestPipeline(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	results := Pipeline(nums)

	if len(results) != len(nums) {
		t.Fatalf("got %d results, want %d", len(results), len(nums))
	}

	// (x * 2) + 10
	want := []int{12, 14, 16, 18, 20}
	for i, v := range results {
		if v != want[i] {
			t.Errorf("results[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestPipeline_Empty(t *testing.T) {
	results := Pipeline([]int{})
	if len(results) != 0 {
		t.Errorf("got %d results, want 0", len(results))
	}
}

func TestRateLimitedFetch(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	urls := make([]string, 5)
	for i := range urls {
		urls[i] = server.URL
	}

	rps := 10
	start := time.Now()
	results := RateLimitedFetch(urls, rps)
	elapsed := time.Since(start)

	if len(results) != 5 {
		t.Fatalf("got %d results, want 5", len(results))
	}

	for i, r := range results {
		if r.Err != nil {
			t.Errorf("results[%d] error: %v", i, r.Err)
		}
		if r.StatusCode != 200 {
			t.Errorf("results[%d] status = %d, want 200", i, r.StatusCode)
		}
	}

	// 5 requests at 10 rps → 최소 ~400ms 이상 걸려야 함
	if elapsed < 400*time.Millisecond {
		t.Errorf("completed too fast (%v), rate limiting may not be working", elapsed)
	}
}

func TestFetchAll_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	urls := []string{server.URL + "/a", server.URL + "/b", server.URL + "/c"}
	results, err := FetchAll(context.Background(), urls)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 3 {
		t.Fatalf("got %d results, want 3", len(results))
	}

	for i, r := range results {
		if r.StatusCode != 200 {
			t.Errorf("results[%d] status = %d, want 200", i, r.StatusCode)
		}
		if !strings.HasSuffix(r.URL, urls[i][len(server.URL):]) {
			t.Errorf("results[%d] URL mismatch", i)
		}
	}
}

func TestFetchAll_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
	}))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	urls := []string{server.URL + "/slow1", server.URL + "/slow2"}
	_, err := FetchAll(ctx, urls)

	if err == nil {
		t.Fatal("expected error from timeout, got nil")
	}
}
