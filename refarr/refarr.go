package refarr

func PassByValue(s []int) {
	for i := range s {
		s[i] = i
	}
}

func PassByReference(s []*int) {
	for i := range s {
		val := i
		s[i] = &val
	}
}
