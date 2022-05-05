package chapter05

import "errors"

func main() {
	a()
}

func a() error {
	return errors.New("some error")
}
