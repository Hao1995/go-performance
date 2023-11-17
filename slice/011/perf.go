package perf

type Person struct {
	Name  string
	Age   int
	Email string
}

func PassByValue(s []Person) {
	for i := range s {
		s[i].Name = "after edit"
		s[i].Age = i
		s[i].Email = "after.edit@example.com"
	}
}

func PassByReference(s []*Person) {
	for i := range s {
		s[i].Name = "after edit"
		s[i].Age = i
		s[i].Email = "after.edit@example.com"
	}
}
