package changesize

import (
	"testing"
)

const (
	DATA_SIZE = 1e6
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

	expected := Person{
		Name:  "after edit",
		Age:   100,
		Email: "after.edit@example.com",
	}

	res := PassByValue(data)
	if item := res[len(res)-1]; item != expected {
		t.Errorf("Expected %+v, got %+v", expected, item)
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

	expected := Person{
		Name:  "after edit",
		Age:   100,
		Email: "after.edit@example.com",
	}

	res := PassByReference(data)
	if item := res[len(res)-1]; *item != expected {
		t.Errorf("Expected %+v, got %+v", expected, *item)
	}

}
