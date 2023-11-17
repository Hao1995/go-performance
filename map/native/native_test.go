package native

import (
	"testing"
)

const DATA_SIZE = 1 << 24

func BenchmarkPassByValue(b *testing.B) {
	data := make(map[int]int, DATA_SIZE)
	for i := 0; i < b.N; i++ {
		PassByValue(data)
	}
}

func TestPassByValue(t *testing.T) {
	data := make(map[int]int, DATA_SIZE)
	PassByValue(data)
	for i := range data {
		if data[i] != i {
			t.Errorf("Expected %d, got %d", i, data[i])
		}
	}
}

func BenchmarkPassByReference(b *testing.B) {
	data := make(map[int]*int, DATA_SIZE)
	for i := 0; i < b.N; i++ {
		PassByReference(data)
	}
}

func TestPassByReference(t *testing.T) {
	data := make(map[int]*int, DATA_SIZE)
	PassByReference(data)
	for i := range data {
		if *data[i] != i {
			t.Errorf("Expected %d, got %d", i, *data[i])
		}
	}
}
