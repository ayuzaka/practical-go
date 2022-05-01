package chapter01

import (
	"fmt"
	"log"
	"strings"
)

func plus(src []string) {
	var title string

	for i, word := range src {
		if i != 0 {
			title += " "
		}
		title += word
	}

	log.Println(title)
}

func builder(src []string) {
	var builder strings.Builder
	builder.Grow(100)

	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}

	log.Println(builder.String())
}

func stringExample() {
	src := []string{"Back", "To", "The", "Future", "Part", "Ⅲ"}
	plus(src)
	builder((src))

	fmt.Println("Hello World")
}
