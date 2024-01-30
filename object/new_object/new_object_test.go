package native

import (
	"testing"
)

func BenchmarkPassByValue(b *testing.B) {
	valPerson := Person{
		Name:  "before edit",
		Age:   -1,
		Email: "before.edit@example.com",
	}
	for i := 0; i < b.N; i++ {
		PassByValue(valPerson)
	}
}

func BenchmarkPassByReference(b *testing.B) {
	ptrPerson := &Person{
		Name:  "before edit",
		Age:   -1,
		Email: "before.edit@example.com",
	}
	for i := 0; i < b.N; i++ {
		PassByReference(ptrPerson)
	}
}
