package native

type Person struct {
	Name  string
	Age   int
	Email string
}

func PassByValue(s Person) Person {
	return Person{
		Name:  "after edit",
		Age:   100,
		Email: "after.edit@example.com",
	}
}

func PassByReference(s *Person) {
	s.Name = "after edit"
	s.Age = 100
	s.Email = "after.edit@example.com"
}
