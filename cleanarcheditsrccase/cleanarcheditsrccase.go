package cleanarcheditsrccase

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

func Value(data []Person) []*ProtoPerson {
	valUsecase(data)
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

func valUsecase(data []Person) {
	for i := range data {
		data[i].Name = fmt.Sprintf("usecase-%d", i)
	}
	valAdapter(data)
}

func valAdapter(data []Person) {
	for i := range data {
		data[i].Name = fmt.Sprintf("adapter-%d", i)
	}
}

func Pointer(data []*Person) []*ProtoPerson {
	ptrUsecase(data)
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

func ptrUsecase(data []*Person) {
	for i := range data {
		data[i].Name = fmt.Sprintf("usecase-%d", i)
	}
	ptrAdapter(data)
}

func ptrAdapter(data []*Person) {
	for i := range data {
		data[i].Name = fmt.Sprintf("adapter-%d", i)
	}
}
