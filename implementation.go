package postfixcli

import (
	"fmt"
	"strconv"
	"strings"
)

func EvaluatePostfix(expression string) (int, error) {
	tokens := strings.Split(expression, " ")
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, fmt.Errorf("недостатньо операндів для операції %s", token)
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("ділення на нуль")
				}
				result = a / b
			}
			stack = append(stack, result)

		default:
			num, err := strconv.Atoi(token)
			if err != nil || num < 0 {
				return 0, fmt.Errorf("некоректний токен: %s", token)
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("некоректний вираз")
	}
	return stack[0], nil
}
