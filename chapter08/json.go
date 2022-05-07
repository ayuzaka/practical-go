package chapter08

import (
	"encoding/json"
	"log"
)

type user struct {
	UserID      string   `json:"user_id"`
	UserName    string   `json:"user_name"`
	Languages   []string `json:"languages"`
	Age         int      `json:"age"`
	CompanyName []string `json:"company_name,omitempty"` // omitempty を指定すると、ゼロ値の場合はエンコード対象外となる
	Foo         *int     `json:"foo,omitempty"`          // ポインタを指定した場合、明示的なゼロ値の場合はエンコードされる
	X           func()   `json:"-"`                      // - を指定することで、エンコード対象外とする
}

func DecodeUser() string {
	u := user{
		UserID:      "001",
		UserName:    "gopher",
		CompanyName: []string{},
		Foo:         Int(0),
		Languages:   []string{}, // 空配列としてエンコードする場合は、空スライスを格納しておく必要がある
	}

	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func Int(v int) *int {
	return &v
}
