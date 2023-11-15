package arrptrmutitime

import "fmt"

type Tmp struct {
	val string
}

func Value(s []Tmp) {
	valueA(s)
	valueB(s)
}

func valueA(s []Tmp) {
	for i := range s {
		s[i].val = fmt.Sprintf("%d-A", i)
	}
}

func valueB(s []Tmp) {
	for i := range s {
		s[i].val = fmt.Sprintf("%d-B", i)
	}
}

func Pointer(s []*Tmp) {
	pointerA(s)
	pointerB(s)
}

func pointerA(s []*Tmp) {
	for i, a := range s {
		a.val = fmt.Sprintf("%d-A", i)
	}
}

func pointerB(s []*Tmp) {
	for i, a := range s {
		a.val = fmt.Sprintf("%d-B", i)
	}
}
