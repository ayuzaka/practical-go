package chapter03

import "fmt"

func String[T any](s T) string {
	return fmt.Sprintf("%v", s)
}

type Struct struct {
	t interface{}
}

func (s Struct) String() string {
	return fmt.Sprintf("%v", s.t)
}

func genericsExample() {
	fmt.Println(String("abc"))
	fmt.Println(String(123))
	fmt.Println(String(true))

	s := Struct{
		t: "Hello World",
	}
	fmt.Println(s.String())
}
