package chapter08

import (
	"reflect"
	"testing"
)

func TestDecodeCSV(t *testing.T) {
	err := DecodeCSV()

	if err != nil {
		t.Errorf("DecodeCSV() error = %v", err)
	}

}

func TestReadCSV(t *testing.T) {
	countries, err := ReadCSV()

	if err != nil {
		t.Errorf("ReadCSV() error = %v", err)
	}

	want := []Country{
		{Name: "アメリカ合衆国", ISOCode: "US/USA", Population: 310232863},
		{Name: "日本", ISOCode: "JP/JPN", Population: 12728000},
		{Name: "中国", ISOCode: "CN/CHN", Population: 1330044000},
	}

	for i, _ := range countries {
		if !reflect.DeepEqual(countries[i], want[i]) {
			t.Errorf("country is mismatch, got=%v, want=%v", countries[i], want[i])
		}
	}
}
