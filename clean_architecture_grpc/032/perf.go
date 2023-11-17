package perf

const (
	DATA_SIZE = 20
)

type InputPerson struct {
	Name  string
	Age   int
	Email string
}

type Person struct {
	Name  string
	Age   int
	Email string
}

type ProtoPerson struct {
	Name  string
	Age   int
	Email string
}

func Value(input []*InputPerson) []*ProtoPerson {
	var people []Person = make([]Person, len(input))
	for i := range input {
		people[i] = Person{
			Name:  "domain-" + input[i].Name,
			Age:   input[i].Age,
			Email: input[i].Email,
		}
	}

	valUsecase(people)

	output := make([]*ProtoPerson, len(people))
	for i := range people {
		output[i] = &ProtoPerson{
			Name:  "proto-" + people[i].Name,
			Age:   people[i].Age,
			Email: people[i].Email,
		}
	}
	return output
}

func valUsecase(data []Person) {
	for i := range data {
		data[i] = Person{
			Name:  "usecase-" + data[i].Name,
			Age:   data[i].Age,
			Email: data[i].Email,
		}
	}
	valAdapter(data)
}

func valAdapter(data []Person) {
	for i := range data {
		data[i] = Person{
			Name:  "adapter-" + data[i].Name,
			Age:   data[i].Age,
			Email: data[i].Email,
		}
	}
}

func Pointer(input []*InputPerson) []*ProtoPerson {
	var people []*Person = make([]*Person, len(input))
	for i := range input {
		people[i] = &Person{
			Name:  "domain-" + input[i].Name,
			Age:   input[i].Age,
			Email: input[i].Email,
		}
	}

	ptrUsecase(people)

	output := make([]*ProtoPerson, len(people))
	for i := range people {
		output[i] = &ProtoPerson{
			Name:  "proto-" + people[i].Name,
			Age:   people[i].Age,
			Email: people[i].Email,
		}
	}

	return output
}

func ptrUsecase(data []*Person) {
	for i := range data {
		data[i] = &Person{
			Name:  "usecase-" + data[i].Name,
			Age:   data[i].Age,
			Email: data[i].Email,
		}
	}
	ptrAdapter(data)
}

func ptrAdapter(data []*Person) {
	for i := range data {
		data[i] = &Person{
			Name:  "adapter-" + data[i].Name,
			Age:   data[i].Age,
			Email: data[i].Email,
		}
	}
}
