package perf

import (
	"testing"
)

const (
	DATA_SIZE = 20
)

func BenchmarkPassByValue(b *testing.B) {
	input := make([]Person, DATA_SIZE)
	for i := range input {
		input[i] = Person{
			Name:  "before edit",
			Age:   -1,
			Email: "before.edit@example.com",
		}
	}
	for i := 0; i < b.N; i++ {
		PassByValue(input)
	}
}

func TestPassByValue(t *testing.T) {
	data := make([]Person, DATA_SIZE)
	for i := range data {
		data[i] = Person{
			Name:  "before edit",
			Age:   -1,
			Email: "before.edit@example.com",
		}
	}

	res := PassByValue(data)
	for i, item := range res {
		expected := Person{
			Name:  "after edit",
			Age:   i,
			Email: "after.edit@example.com",
		}
		if item != expected {
			t.Errorf("Expected %+v, got %+v", expected, item)
		}
	}
}

func BenchmarkPassByReference(b *testing.B) {
	data := make([]*Person, DATA_SIZE)
	for i := range data {
		data[i] = &Person{
			Name:  "before edit",
			Age:   -1,
			Email: "before.edit@example.com",
		}
	}
	for i := 0; i < b.N; i++ {
		PassByReference(data)
	}
}

func TestPassByReference(t *testing.T) {
	data := make([]*Person, DATA_SIZE)
	for i := range data {
		data[i] = &Person{
			Name:  "before edit",
			Age:   -1,
			Email: "before.edit@example.com",
		}
	}
	PassByReference(data)
	for i := range data {
		if data[i].Age != i {
			t.Errorf("Expected %d, got %d", i, data[i].Age)
		}
	}
}
