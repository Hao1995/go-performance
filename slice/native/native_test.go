package native

import (
	"testing"
)

const (
	DATA_SIZE = 1e6
)

func BenchmarkPassByValue(b *testing.B) {
	valPeople := make([]Person, DATA_SIZE)
	for i := range valPeople {
		valPeople[i] = Person{
			Name:  "before edit",
			Age:   -1,
			Email: "before.edit@example.com",
		}
	}
	for i := 0; i < b.N; i++ {
		PassByValue(valPeople)
	}
}

func TestPassByValue(t *testing.T) {
	arr := make([]Person, 5)
	for i := range arr {
		arr[i].Age = -1
	}

	PassByValue(arr)
	for i := range arr {
		if arr[i].Age != i {
			t.Errorf("Expected %d, got %d", i, arr[i].Age)
		}
	}
}

func BenchmarkPassByReference(b *testing.B) {
	ptrPeople := make([]*Person, DATA_SIZE)
	for i := range ptrPeople {
		ptrPeople[i] = &Person{
			Name:  "before edit",
			Age:   -1,
			Email: "before.edit@example.com",
		}
	}

	for i := 0; i < b.N; i++ {
		PassByReference(ptrPeople)
	}
}

func TestPassByReference(t *testing.T) {
	arr := make([]*Person, 5)
	for i := range arr {
		arr[i] = &Person{Age: -1}
	}

	PassByReference(arr)
	for i := range arr {
		if arr[i].Age != i {
			t.Errorf("Expected %d, got %d", i, arr[i].Age)
		}
	}
}
