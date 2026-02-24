package channel

// sumWithChannel은 nums를 앞/뒤 절반으로 나누어
// 각각 goroutine에서 합산한 뒤, channel을 통해 결과를 모아 총합을 반환한다.
func sumWithChannel(nums []int) int {
	// TODO: 구현하세요
	return 0
}

// generateSequence는 0부터 n-1까지의 정수를 순서대로 보내는 receive-only channel을 반환한다.
// 모든 값을 보낸 후 channel을 닫아야 한다.
func generateSequence(n int) <-chan int {
	// TODO: 구현하세요
	return nil
}

// merge는 두 개의 receive-only channel을 하나로 합친다.
// 두 입력 channel이 모두 닫히면 출력 channel도 닫혀야 한다.
// select를 사용하여 구현할 것.
func merge(ch1, ch2 <-chan int) <-chan int {
	// TODO: 구현하세요
	return nil
}
