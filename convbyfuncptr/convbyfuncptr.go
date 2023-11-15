package convbyfuncptr

type Tmp struct {
	val int
}

func PassByValue(s []int) []Tmp {
	res := make([]Tmp, len(s))
	for i := range s {
		v := genTmp(i)
		res[i] = *v
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
