package cleanarcheditsrccase

import (
	"fmt"
	"testing"
)

var (
	valueData   []Person
	pointerData []*Person
)

func init() {
	valueData = make([]Person, DATA_SIZE)
	for i := 0; i < DATA_SIZE; i++ {
		valueData[i] = Person{Name: fmt.Sprintf("init-%d", i)}
	}

	pointerData = make([]*Person, DATA_SIZE)
	for i := 0; i < DATA_SIZE; i++ {
		pointerData[i] = &Person{Name: fmt.Sprintf("init-%d", i)}
	}
}

func BenchmarkValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Value(valueData)
	}
}

func BenchmarkValueAdapter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valAdapter(valueData)
	}
}

func BenchmarkValueUsecase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valUsecase(valueData)
	}
}

func BenchmarkValueHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valHandler(valueData)
	}
}

func TestValue(t *testing.T) {
	res := Value(valueData)
	for i, item := range res {
		if item.Name != fmt.Sprintf("proto-adapter-%d", i) {
			t.Errorf("Expected %s, got %s", fmt.Sprintf("proto-adapter-%d", i), item.Name)
		}
	}
}

func BenchmarkPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pointer(pointerData)
	}
}

func BenchmarkPointerAdapter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptrAdapter(pointerData)
	}
}

func BenchmarkPointerUsecase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptrUsecase(pointerData)
	}
}

func BenchmarkPointerHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptrHandler(pointerData)
	}
}

func TestPointer(t *testing.T) {
	res := Pointer(pointerData)
	for i, item := range res {
		if item.Name != fmt.Sprintf("proto-adapter-%d", i) {
			t.Errorf("Expected %s, got %s", fmt.Sprintf("proto-adapter-%d", i), item.Name)
		}
	}
}
