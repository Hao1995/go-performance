package convtoarrptr

type Tmp struct {
	val int
}

func PassByValue(s []int) []*Tmp {
	res := make([]*Tmp, len(s))
	for i := range s {
		res[i] = genTmp(i)
	}
	return res
}

func PassByReference(s []*int) []*Tmp {
	res := make([]*Tmp, len(s))
	for i := range s {
		res[i] = genTmp(i)
	}
	return res
}

func genTmp(i int) *Tmp {
	return &Tmp{val: i}
}
