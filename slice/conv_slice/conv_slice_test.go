package convslice

import (
	"testing"
)

const (
	DATA_SIZE = 1e6
)

func BenchmarkPassByValue(b *testing.B) {
	valPeople := make([]People, DATA_SIZE)
	for i := range valPeople {
		valPeople[i] = People{
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
	arr := make([]People, 5)
	for i := range arr {
		arr[i] = People{
			Name:  "before edit",
			Age:   -1,
			Email: "before.edit@example.com",
		}
	}

	expected := Member{
		Name:  "after edit",
		Age:   100,
		Email: "after.edit@example.com",
	}
	res := PassByValue(arr)
	for i := range arr {
		if res[i] != expected {
			t.Errorf("Expected %+v, got %+v", expected, res[i])
		}
	}
}

func BenchmarkPassByReference(b *testing.B) {
	ptrPeople := make([]*People, DATA_SIZE)
	for i := range ptrPeople {
		ptrPeople[i] = &People{
			Name:  "before edit",
			Age:   -1,
			Email: "after.edit@example.com",
		}
	}
	for i := 0; i < b.N; i++ {
		PassByReference(ptrPeople)
	}
}

func TestPassByReference(t *testing.T) {
	arr := make([]*People, 5)
	for i := range arr {
		arr[i] = &People{
			Name:  "before edit",
			Age:   -1,
			Email: "before.edit@example.com",
		}
	}

	expected := Member{
		Name:  "after edit",
		Age:   100,
		Email: "after.edit@example.com",
	}

	res := PassByReference(arr)
	for i := range arr {
		if *res[i] != expected {
			t.Errorf("Expected %+v, got %+v", expected, *res[i])
		}
	}
}
