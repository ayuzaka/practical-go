package chapter02

import (
	"encoding/json"
	"fmt"
)

type Stringer interface {
	String() string
}

type GoStringer interface {
	GoString() string
}

type ConfidentialCustomer struct {
	CustomerID int64
	CreditCard CreditCard
}

type CreditCard string

func (c CreditCard) String() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func (c CreditCard) GoString() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func secretExample() {
	c := ConfidentialCustomer{
		CustomerID: 1,
		CreditCard: "4111-1111-1111-1111",
	}

	fmt.Println(c)
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%#v\n", c)

	bytes, _ := json.Marshal(c)
	fmt.Println("JSON: ", string(bytes))
}
