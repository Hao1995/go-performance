package mapcase

func PassByValue(s map[int]int) {
	res := make(map[int]int, len(s))
	for i := range s {
		res[i] = i
	}
}

func PassByReference(s map[int]*int) {
	res := make(map[int]*int, len(s))
	for i := range s {
		j := new(int)
		*j = i
		res[i] = j
	}
}
