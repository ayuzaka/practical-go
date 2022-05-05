package chapter04

import "fmt"

func sliceCastExample() {
	var fishList = []any{"鯖", "鰤", "鮪"}
	fishNames := make([]string, len(fishList))
	for i, f := range fishList {
		// ダウンキャストは型アサーションが必要
		if fn, ok := f.(string); ok {
			fishNames[i] = fn
		}
	}

	fmt.Println(fishNames)

	fibonacciNumbers := []int{1, 1, 2, 3, 5, 8}
	anyValues := make([]any, len(fibonacciNumbers))
	for i, fn := range fibonacciNumbers {
		anyValues[i] = fn
	}
}
