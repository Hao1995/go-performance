package convtoarrptr

import (
	"testing"
)

func BenchmarkPassByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 1)
		PassByValue(arr)
	}
}

func TestPassByValue(t *testing.T) {
	arr := make([]int, 5)
	res := PassByValue(arr)
	for i := range arr {
		if res[i].val != i {
			t.Errorf("Expected %d, got %d", i, res[i].val)
		}
	}
}

func BenchmarkPassByReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]*int, 1)
		PassByReference(arr)
	}
}

func TestPassByReference(t *testing.T) {
	arr := make([]*int, 5)
	res := PassByReference(arr)
	for i := range arr {
		if res[i].val != i {
			t.Errorf("Expected %d, got %d", i, res[i].val)
		}
	}
}
