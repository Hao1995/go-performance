package getslice

import (
	"fmt"
	"testing"
)

var (
	valData = valAdapter()
	ptrData = ptrAdapter()
)

func BenchmarkValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Value()
	}
}

func BenchmarkValueAdapter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valAdapter()
	}
}

func BenchmarkValueUsecase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valUsecase()
	}
}

func BenchmarkValueHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valHandler(valData)
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

func BenchmarkPointerAdapter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptrAdapter()
	}
}

func BenchmarkPointerUsecase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptrUsecase()
	}
}

func BenchmarkPointerHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptrHandler(ptrData)
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
