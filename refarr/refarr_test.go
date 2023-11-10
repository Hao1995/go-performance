package refarr

import (
	"testing"
)

func BenchmarkPassByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 1<<24)
		PassByValue(arr)
	}
}

func TestPassByValue(t *testing.T) {
	arr := make([]int, 5)
	PassByValue(arr)
	for i := range arr {
		if arr[i] != i {
			t.Errorf("Expected %d, got %d", i, arr[i])
		}
	}
}

func BenchmarkPassByReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]*int, 1<<24)
		PassByReference(arr)
	}
}

func TestPassByReference(t *testing.T) {
	arr := make([]*int, 5)
	PassByReference(arr)
	for i := range arr {
		if *arr[i] != i {
			t.Errorf("Expected %d, got %d", i, *arr[i])
		}
	}
}
