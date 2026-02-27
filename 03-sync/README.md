# 03. Sync Primitives

## 목표

`sync` 패키지의 동기화 프리미티브를 이해하고, channel 대신 이것들을 써야 하는 상황을 구분한다.

## 과제

`sync.go`의 함수 3개를 구현하세요.

### 1. `SafeCounter` 구조체
- 여러 goroutine에서 동시에 `Increment()`와 `Value()`를 호출할 수 있는 카운터
- `sync.Mutex`를 사용하여 동시 접근을 보호할 것

### 2. `FetchOnce` 구조체
- `Get()` 메서드가 여러 goroutine에서 동시에 호출되더라도 실제 데이터 로딩은 **딱 한 번만** 실행
- `sync.Once`를 사용할 것
- 로딩 함수는 생성 시 주입받음

### 3. `ParallelSum(nums []int, workers int) int`
- slice를 workers 개수만큼 분할하여 병렬로 합산
- `sync.WaitGroup`으로 모든 worker 완료를 대기
- 각 worker의 부분합을 안전하게 합산할 것 (`sync.Mutex` 또는 `atomic` 사용)

## 실행

```bash
cd 03-sync
go test -v -race
```
