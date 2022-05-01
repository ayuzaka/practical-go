package chapter02

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func typeExample() {
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))

	fmt.Println(fileType.Implements(writerType))
}
