package channel

import (
	"sort"
	"testing"
)

func TestSumWithChannel(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"even", []int{1, 2, 3, 4}, 10},
		{"odd", []int{1, 2, 3, 4, 5}, 15},
		{"negative", []int{-1, -2, 3, 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumWithChannel(tt.nums)
			if got != tt.want {
				t.Errorf("sumWithChannel(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestGenerateSequence(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"zero", 0},
		{"one", 1},
		{"five", 5},
		{"ten", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := generateSequence(tt.n)
			if ch == nil {
				t.Fatal("generateSequence returned nil channel")
			}

			var got []int
			for v := range ch {
				got = append(got, v)
			}

			if len(got) != tt.n {
				t.Errorf("got %d values, want %d", len(got), tt.n)
			}

			for i, v := range got {
				if v != i {
					t.Errorf("got[%d] = %d, want %d", i, v, i)
				}
			}
		})
	}
}

func TestMerge(t *testing.T) {
	ch1 := generateSequence(5)  // 0,1,2,3,4
	ch2 := generateSequence(3)  // 0,1,2

	merged := merge(ch1, ch2)
	if merged == nil {
		t.Fatal("merge returned nil channel")
	}

	var got []int
	for v := range merged {
		got = append(got, v)
	}

	if len(got) != 8 {
		t.Fatalf("got %d values, want 8", len(got))
	}

	// merge 순서는 비결정적이므로 정렬 후 비교
	sort.Ints(got)
	want := []int{0, 0, 1, 1, 2, 2, 3, 4}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("sorted got[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}
