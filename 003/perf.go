package perf

type People struct {
	Name  string
	Age   int
	Email string
}

type Member struct {
	Name  string
	Age   int
	Email string
}

func PassByValue(s []People) []Member {
	res := make([]Member, len(s))
	for i := range s {
		res[i] = Member{
			Name:  "after edit",
			Age:   100,
			Email: "after.edit@example.com",
		}
	}
	return res
}

func PassByReference(s []*People) []*Member {
	res := make([]*Member, len(s))
	for i := range s {
		res[i] = &Member{
			Name:  "after edit",
			Age:   100,
			Email: "after.edit@example.com",
		}
	}
	return res
}
