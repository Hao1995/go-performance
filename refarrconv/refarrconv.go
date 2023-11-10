package refarrconv

type Tmp struct {
	val int
}

func PassByValue(s []int) []Tmp {
	res := make([]Tmp, len(s))
	for i := range s {
		res[i] = Tmp{val: i}
	}
	return res
}

func PassByReference(s []*int) []*Tmp {
	res := make([]*Tmp, len(s))
	for i := range s {
		res[i] = &Tmp{val: i}
	}
	return res
}
