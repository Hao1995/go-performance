package cleanarchcas

import (
	"fmt"
	"testing"
)

func BenchmarkValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Value()
	}
}

func TestValue(t *testing.T) {
	res := Value()
	for i, item := range res {
		if item.Name != fmt.Sprintf("proto-usecase-%d", i) {
			t.Errorf("Expected %s, got %s", fmt.Sprintf("proto-usecase-%d", i), item.Name)
		}
	}
}

func BenchmarkPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pointer()
	}
}

func TestPointer(t *testing.T) {
	res := Pointer()
	for i, item := range res {
		if item.Name != fmt.Sprintf("proto-usecase-%d", i) {
			t.Errorf("Expected %s, got %s", fmt.Sprintf("proto-usecase-%d", i), item.Name)
		}
	}
}
