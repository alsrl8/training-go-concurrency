# 01. Channel Basics

## 목표

channel의 기본 동작을 이해하고, unbuffered/buffered channel의 차이를 체감한다.

## 과제

`channel.go`의 함수 3개를 구현하세요.

### 1. `sumWithChannel(nums []int) int`
- goroutine을 사용해 슬라이스를 **앞쪽 절반 / 뒤쪽 절반**으로 나누어 합산
- 각 goroutine이 channel로 부분합을 보내고, 메인에서 합산하여 반환

### 2. `generateSequence(n int) <-chan int`
- 0부터 n-1까지의 정수를 channel로 순서대로 보내는 goroutine을 실행
- **receive-only channel**을 반환
- 모든 값을 보낸 후 channel을 닫을 것

### 3. `merge(ch1, ch2 <-chan int) <-chan int`
- 두 개의 channel로부터 값을 받아 하나의 channel로 합쳐서 반환
- 두 입력 channel이 모두 닫히면 출력 channel도 닫힐 것
- `select`를 사용할 것

## 실행

```bash
cd 01-channel-basics
go test -v -race
```

모든 테스트가 통과하면 성공입니다.
