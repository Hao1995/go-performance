package createslice

import (
	"fmt"
	"testing"
)

func BenchmarkValue(b *testing.B) {
	var input []*InputPerson = make([]*InputPerson, DATA_SIZE)
	for i := 0; i < DATA_SIZE; i++ {
		input[i] = &InputPerson{
			Name:  fmt.Sprintf("init-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
	}

	for i := 0; i < b.N; i++ {
		Value(input)
	}
}

func BenchmarkValueUsecase(b *testing.B) {
	data := make([]Person, DATA_SIZE)
	for i := range data {
		data[i] = Person{
			Name:  fmt.Sprintf("test-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
	}
	for i := 0; i < b.N; i++ {
		valUsecase(data)
	}
}

func BenchmarkValueAdapter(b *testing.B) {
	data := make([]Person, DATA_SIZE)
	for i := range data {
		data[i] = Person{
			Name:  fmt.Sprintf("test-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
	}
	for i := 0; i < b.N; i++ {
		valAdapter(data)
	}
}

func TestValue(t *testing.T) {
	input := make([]*InputPerson, DATA_SIZE)
	for i := range input {
		input[i] = &InputPerson{
			Name:  fmt.Sprintf("input-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
	}

	res := Value(input)
	for i, item := range res {
		expected := ProtoPerson{
			Name:  fmt.Sprintf("proto-adapter-usecase-domain-input-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
		if *item != expected {
			t.Errorf("Expected %+v, got %+v", expected, *item)
		}
	}
}

func BenchmarkPointer(b *testing.B) {
	input := make([]*InputPerson, DATA_SIZE)
	for i := range input {
		input[i] = &InputPerson{
			Name:  fmt.Sprintf("input-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
	}
	for i := 0; i < b.N; i++ {
		Pointer(input)
	}
}

func BenchmarkPointerUsecase(b *testing.B) {
	data := make([]*Person, DATA_SIZE)
	for i := range data {
		data[i] = &Person{
			Name:  fmt.Sprintf("test-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
	}
	for i := 0; i < b.N; i++ {
		ptrUsecase(data)
	}
}

func BenchmarkPointerAdapter(b *testing.B) {
	data := make([]*Person, DATA_SIZE)
	for i := range data {
		data[i] = &Person{
			Name:  fmt.Sprintf("test-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
	}
	for i := 0; i < b.N; i++ {
		ptrAdapter(data)
	}
}

func TestPointer(t *testing.T) {
	input := make([]*InputPerson, DATA_SIZE)
	for i := range input {
		input[i] = &InputPerson{
			Name:  fmt.Sprintf("input-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
	}
	res := Pointer(input)
	for i, item := range res {
		expected := ProtoPerson{
			Name:  fmt.Sprintf("proto-adapter-usecase-domain-input-%d", i),
			Age:   -1,
			Email: "test@example.com",
		}
		if *item != expected {
			t.Errorf("Expected %+v, got %+v", expected, *item)
		}
	}
}
