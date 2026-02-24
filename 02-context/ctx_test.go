package ctx

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"
)

func TestFetchWithTimeout_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	code, err := fetchWithTimeout(server.URL, 2*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if code != 200 {
		t.Errorf("got status %d, want 200", code)
	}
}

func TestFetchWithTimeout_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	_, err := fetchWithTimeout(server.URL, 100*time.Millisecond)
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
}

func TestDoWork_Cancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 350*time.Millisecond)
	defer cancel()

	results := make(chan int, 100)
	go doWork(ctx, results)

	// 350ms timeout, 100ms interval → 최대 3개 정도 기대
	var got []int
	for v := range results {
		got = append(got, v)
	}

	if len(got) == 0 {
		t.Fatal("doWork produced no results")
	}
	if len(got) > 5 {
		t.Errorf("doWork produced too many results (%d), cancellation may not be working", len(got))
	}

	// 값이 순차적인지 확인
	for i, v := range got {
		if v != i+1 {
			t.Errorf("got[%d] = %d, want %d", i, v, i+1)
		}
	}
}

func TestFanOutWithCancel_AllComplete(t *testing.T) {
	ctx := context.Background()
	results := fanOutWithCancel(ctx, 5)

	sort.Ints(results)
	if len(results) != 5 {
		t.Fatalf("got %d results, want 5", len(results))
	}
	for i, v := range results {
		if v != i {
			t.Errorf("results[%d] = %d, want %d", i, v, i)
		}
	}
}

func TestFanOutWithCancel_Cancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // 즉시 취소

	results := fanOutWithCancel(ctx, 100)

	// 즉시 취소했으므로 결과가 매우 적거나 없어야 함
	if len(results) > 50 {
		t.Errorf("got %d results after immediate cancel, expected much fewer", len(results))
	}
}
