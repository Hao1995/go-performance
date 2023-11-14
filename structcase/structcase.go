package structcase

type Person struct {
	Name  string
	Age   int
	Email string
}

func PassByValue(s Person) {
	s.Name = "John"
	s.Age = 1
	s.Email = "test@example.com"
}

func PassByReference(s *Person) {
	s.Name = "John"
	s.Age = 1
	s.Email = "test@example.com"
}
