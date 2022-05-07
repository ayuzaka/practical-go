package chapter08

import "testing"

func TestDecodeUser(t *testing.T) {
	want := `{"user_id":"001","user_name":"gopher","languages":[],"age":0,"foo":0}`
	got := DecodeUser()

	if got != want {
		t.Errorf("DelayDecode() = %v, want %v", got, want)
	}

}
