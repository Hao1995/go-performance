package arrptrmutitime

import (
	"fmt"
	"testing"
)

const (
	DATA_SIZE = 10000
)

var (
	valData = make([]Tmp, DATA_SIZE)
	ptrData = make([]*Tmp, DATA_SIZE)
)

func init() {
	for i := range valData {
		valData[i] = Tmp{val: "init"}
	}
	for i := range ptrData {
		ptrData[i] = &Tmp{val: "init"}
	}
}

func BenchmarkValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Value(valData)
	}
}

func TestValue(t *testing.T) {
	Value(valData)
	for i, a := range valData {
		if a.val != fmt.Sprintf("%d-B", i) {
			t.Errorf("Expected `%d-B`, got %v", i, a.val)
		}
	}
}

func BenchmarkPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pointer(ptrData)
	}
}

func TestPointer(t *testing.T) {
	Pointer(ptrData)
	for i, a := range ptrData {
		if a.val != fmt.Sprintf("%d-B", i) {
			t.Errorf("Expected `%d-B`, got %v", i, a.val)
		}
	}
}
