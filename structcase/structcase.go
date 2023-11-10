package structcase

type LargeStruct struct {
	Data [1 << 24]int // 16MB of integers
}

func PassByValue(s LargeStruct) {
	for i := range s.Data {
		s.Data[i] = i
	}
}

func PassByReference(s *LargeStruct) {
	for i := range s.Data {
		s.Data[i] = i
	}
}
