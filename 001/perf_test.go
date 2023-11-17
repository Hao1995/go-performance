package perf

import (
	"testing"
)

var (
	valPerson = Person{
		Name:  "before edit",
		Age:   -1,
		Email: "before.edit@example.com",
	}

	ptrPerson = &Person{
		Name:  "before edit",
		Age:   -1,
		Email: "before.edit@example.com",
	}
)

func BenchmarkPassByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassByValue(valPerson)
	}
}

func TestPassByValue(t *testing.T) {
	s := Person{
		Name:  "before edit",
		Age:   -1,
		Email: "before.edit@example.com",
	}

	expected := Person{
		Name:  "after edit",
		Age:   100,
		Email: "after.edit@example.com",
	}

	res := PassByValue(s)
	if res != expected {
		t.Errorf("Expected %+v, got %+v", expected, res)
	}
}

func BenchmarkPassByReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassByReference(ptrPerson)
	}
}

func TestPassByReference(t *testing.T) {
	s := &Person{
		Name:  "before edit",
		Age:   -1,
		Email: "before.edit@example.com",
	}

	expected := &Person{
		Name:  "after edit",
		Age:   100,
		Email: "after.edit@example.com",
	}

	PassByReference(s)

	if *s != *expected {
		t.Errorf("Expected %+v, got %+v", expected, s)
	}
}
