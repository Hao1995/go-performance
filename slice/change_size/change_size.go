package changesize

type Person struct {
	Name  string
	Age   int
	Email string
}

func PassByValue(s []Person) []Person {
	s = append(s, Person{
		Name:  "after edit",
		Age:   100,
		Email: "after.edit@example.com",
	})
	return s
}

func PassByReference(s []*Person) []*Person {
	s = append(s, &Person{
		Name:  "after edit",
		Age:   100,
		Email: "after.edit@example.com",
	})
	return s
}
