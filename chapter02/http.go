package chapter02

import "fmt"

type HTTPStatus int

const (
	StatusOK              HTTPStatus = 200
	StatusUnauthorized    HTTPStatus = 401
	StatusPaymentRequired HTTPStatus = 402
	StatusForbidden       HTTPStatus = 403
)

func (s HTTPStatus) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	default:
		return fmt.Sprintf("HTTPStatus (%d)", s)
	}
}

func httpExample() {
	var status1 HTTPStatus = 200
	fmt.Println(status1.String())

	var status2 HTTPStatus = 402
	fmt.Println(status2.String())

	var status3 HTTPStatus = 500
	fmt.Println(status3.String())
}
