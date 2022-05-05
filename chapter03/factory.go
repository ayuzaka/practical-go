package chapter03

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

// ファクトリー関数
func NewPerson(first, last string) *Person {
	return &Person{
		FirstName: first,
		LastName:  last,
	}
}

func (p Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// 状態を変更する場合は、ポインターにする
func (p *Person) UpdateFirstName(name string) {
	p.FirstName = name
}

func factoryExample() {
	person1 := NewPerson("John", "Lennon")
	fmt.Printf("%v\n", person1)

	person2 := &Person{
		FirstName: "Paul",
		LastName:  "McCartney",
	}
	fmt.Printf("%v\n", person2)

	var person3 Person
	fmt.Printf("%v\n", person3)

	fmt.Println(person1.GetFullName())

	person1.UpdateFirstName("Ringo")
	fmt.Println(person1.GetFullName())
}
