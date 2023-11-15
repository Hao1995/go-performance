package structcase

import (
	"testing"
)

func BenchmarkPassByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s Person = Person{}
		PassByValue(s)
	}
}

func TestPassByValue(t *testing.T) {
	s := Person{
		Name: "init",
	}
	newObj := PassByValue(s)
	if newObj.Name != "John" {
		t.Errorf("Expected %s, got %s", "", s.Name)
	}
}

func BenchmarkPassByReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s *Person = &Person{}
		PassByReference(s)
	}
}

func TestPassByReference(t *testing.T) {
	s := &Person{
		Name: "init",
	}
	PassByReference(s)
	if s.Name != "John" {
		t.Errorf("Expected %s, got %s", "", s.Name)
	}
}
