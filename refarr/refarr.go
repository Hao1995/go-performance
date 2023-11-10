package refarr

func PassByValue(s []int) {
	for i := range s {
		s[i] = i
	}
}

func PassByReference(s []*int) {
	for i := range s {
		j := new(int)
		*j = i
		s[i] = j
	}
}
