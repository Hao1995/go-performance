package mapcase

import (
	"testing"
)

func BenchmarkPassByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapData := make(map[int]int, 1<<24)
		PassByValue(mapData)
	}
}

func TestPassByValue(t *testing.T) {
	mapData := make(map[int]int, 1<<24)
	PassByValue(mapData)
	for i := range mapData {
		if mapData[i] != i {
			t.Errorf("Expected %d, got %d", i, mapData[i])
		}
	}
}

func BenchmarkPassByReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapData := make(map[int]*int, 1<<24)
		PassByReference(mapData)
	}
}

func TestPassByReference(t *testing.T) {
	mapData := make(map[int]*int, 1<<24)
	PassByReference(mapData)
	for i := range mapData {
		if *mapData[i] != i {
			t.Errorf("Expected %d, got %d", i, *mapData[i])
		}
	}
}
