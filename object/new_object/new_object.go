package native

type Person struct {
	Name  string
	Age   int
	Email string
}

type Member struct {
	MName  string
	MAge   int
	MEmail string
}

func PassByValue(s Person) Member {
	return Member{
		MName:  s.Name,
		MAge:   s.Age,
		MEmail: s.Email,
	}
}

func PassByReference(s *Person) *Member {
	return &Member{
		MName:  s.Name,
		MAge:   s.Age,
		MEmail: s.Email,
	}
}
