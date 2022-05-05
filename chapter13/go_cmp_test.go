package chapter13

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestByCmp(t *testing.T) {
	tom := User{UserID: "0001", UserName: "Tom", Languages: []string{"Java", "Go"}}
	tom2 := User{UserID: "0001", UserName: "Tom", Languages: []string{"Java", "Go"}}

	if diff := cmp.Diff(tom, tom2); diff != "" {
		t.Errorf("User tom is mismatch, tom=%v, tom2=%v", tom, tom2)
	}
}

func TestAllowUnexports(t *testing.T) {
	type X struct {
		numUnExport int
		NumExport   int
	}

	num1 := X{100, -1}
	num2 := X{100, -1}

	opt := cmp.AllowUnexported(X{})

	if diff := cmp.Diff(num1, num2, opt); diff != "" {
		t.Errorf("X value is mismatch (-num +num2):%s\n", diff)
	}
}

func TestIgnoreUnexported(t *testing.T) {
	type X struct {
		numUnExport int
		NumExport   int
	}

	num1 := X{100, -1}
	num2 := X{999, -1}

	opt := cmpopts.IgnoreUnexported(X{})

	if diff := cmp.Diff(num1, num2, opt); diff != "" {
		t.Errorf("X value is mismatch (-num +num2):%s\n", diff)
	}
}

func TestIgnoreFields(t *testing.T) {
	type X struct {
		NumExport int
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	num1 := X{-1, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}
	num2 := X{-1, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}

	opt := cmpopts.IgnoreFields(X{}, "CreatedAt", "UpdatedAt")

	if diff := cmp.Diff(num1, num2, opt); diff != "" {
		t.Errorf("X value is mismatch (-num +num2):%s\n", diff)
	}
}
