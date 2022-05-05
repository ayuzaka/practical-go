package chapter03

import "fmt"

type Book struct {
	Title string
	ISBN  string
}

func (b Book) GetAmazonURL() string {
	return "https://amazon.co.jp/dp/" + b.ISBN
}

type OreillyBook struct {
	Book
	ISBN13 string
}

func (o OreillyBook) GetOreillyURL() string {
	return "https://www.oreilly.co.jp/books" + o.ISBN13 + "/"
}

func embedExample() {
	ob := OreillyBook{
		ISBN13: "9784873119038",
		Book: Book{
			Title: "Real World HTTP",
			ISBN:  "9704773119039",
		},
	}

	fmt.Println(ob.GetAmazonURL())
	fmt.Println(ob.GetOreillyURL())
}
