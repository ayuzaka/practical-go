package chapter13

import "fmt"

func Calc(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("0除算は未定義です")
		}

		return a / b, nil
	}

	return 0, fmt.Errorf("予期しない演算子 %v が指定されました", operator)
}
