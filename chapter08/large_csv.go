package chapter08

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type record struct {
	Number  int    `csv:"number"`
	Message string `csv:"message"`
}

func Write() error {
	c := make(chan interface{})
	go func() {
		defer close(c)

		for i := 0; i < 1000*1000; i++ {
			c <- record{
				Message: "Hello",
				Number:  i + 1,
			}
		}

		return
	}()

	if err := gocsv.MarshalChan(c, gocsv.DefaultCSVWriter(os.Stdout)); err != nil {
		return err
	}

	return nil
}

func Read() {
	f, err := os.Open("large.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	c := make(chan record)
	done := make(chan bool)
	go func() {
		if err := gocsv.UnmarshalToChan(f, c); err != nil {
			log.Fatal(err)
		}
		done <- true
	}()

	for {
		select {
		case v := <-c:
			fmt.Printf("%+v\n", v)
		case <-done:
			return
		}
	}

}
