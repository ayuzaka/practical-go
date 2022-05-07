package chapter08

import "testing"

func TestWrite(t *testing.T) {
	err := Write()

	if err != nil {
		t.Errorf("Write() error = %v", err)
	}
}
