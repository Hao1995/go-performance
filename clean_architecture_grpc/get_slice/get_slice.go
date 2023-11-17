package getslice

import "fmt"

const (
	DATA_SIZE = 20
)

type Person struct {
	Name string
}

type ProtoPerson struct {
	Name string
}

func Value() []*ProtoPerson {
	data := valUsecase()
	return valHandler(data)
}

func valHandler(data []Person) []*ProtoPerson {
	protoData := make([]*ProtoPerson, len(data))
	for i := range data {
		protoData[i] = &ProtoPerson{
			Name: fmt.Sprintf("proto-%s", data[i].Name),
		}
	}
	return protoData
}

func valUsecase() []Person {
	data := valAdapter()
	for i := range data {
		data[i].Name = fmt.Sprintf("usecase-%d", i)
	}
	return data
}

func valAdapter() []Person {
	data := make([]Person, DATA_SIZE)
	for i := 0; i < DATA_SIZE; i++ {
		data[i] = Person{Name: fmt.Sprintf("adapter-%d", i)}
	}
	return data
}

func Pointer() []*ProtoPerson {
	data := ptrUsecase()
	return ptrHandler(data)
}

func ptrHandler(data []*Person) []*ProtoPerson {
	protoData := make([]*ProtoPerson, len(data))
	for i := range data {
		protoData[i] = &ProtoPerson{
			Name: fmt.Sprintf("proto-%s", data[i].Name),
		}
	}
	return protoData
}

func ptrUsecase() []*Person {
	data := ptrAdapter()
	for i := range data {
		data[i].Name = fmt.Sprintf("usecase-%d", i)
	}
	return data
}

func ptrAdapter() []*Person {
	data := make([]*Person, DATA_SIZE)
	for i := 0; i < DATA_SIZE; i++ {
		data[i] = &Person{Name: fmt.Sprintf("adapter-%d", i)}
	}
	return data
}
