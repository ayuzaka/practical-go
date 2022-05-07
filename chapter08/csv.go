package chapter08

import (
	"os"

	"github.com/gocarina/gocsv"
)

type Country struct {
	Name       string `csv:"国名"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人口"`
}

func DecodeCSV() error {
	lines := []Country{
		{Name: "アメリカ合衆国", ISOCode: "US/USA", Population: 310232863},
		{Name: "日本", ISOCode: "JP/JPN", Population: 12728000},
		{Name: "中国", ISOCode: "CN/CHN", Population: 1330044000},
	}

	f, err := os.Create("country.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	if err := gocsv.MarshalFile(&lines, f); err != nil {
		return err
	}

	return nil
}

func ReadCSV() ([]Country, error) {
	f, err := os.Open("country.csv")
	if err != nil {
		return []Country{}, err
	}
	defer f.Close()

	var lines []Country
	if err := gocsv.UnmarshalFile(f, &lines); err != nil {
		return []Country{}, err
	}

	return lines, nil
}
