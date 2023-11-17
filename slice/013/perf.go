package perf

type Person struct {
	Name  string
	Age   int
	Email string
}

func PassByValue(s []Person) []Person {
	res := make([]Person, len(s))
	for i := range s {
		res[i] = Person{
			Name:  "after edit",
			Age:   i,
			Email: "after.edit@example.com",
		}
	}
	return res
}

func PassByReference(s []*Person) {
	for i := range s {
		s[i] = &Person{
			Name:  "after edit",
			Age:   i,
			Email: "after.edit@example.com",
		}
	}
}
