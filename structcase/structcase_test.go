package structcase

import (
	"testing"
)

func BenchmarkPassByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		PassByValue(s)
	}
}

func BenchmarkPassByReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		PassByReference(&s)
	}
}
