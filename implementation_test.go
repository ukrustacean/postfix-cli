package postfixcli

import (
	"fmt"
	"testing"
)

func TestEvaluatePostfix(t *testing.T) {
	res, err := EvaluatePostfix("4 2 - 3 * 5 +")

	if err != nil {
		t.Error("Error must be nil")
	}

	if res != 11 {
		t.Error("Postfix expression evaluated incorrectly")
	}

}

func ExampleEvaluatePostfix() {
	res, _ := EvaluatePostfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 4
}

func TestEvaluateSimpleAddition(t *testing.T) {
	expr := "2 3 +"
	expected := 5

	result, err := EvaluatePostfix(expr)
	if err != nil {
		t.Fatalf("Неочікувана помилка: %v", err)
	}
	if result != expected {
		t.Errorf("Очікувалось %d, отримано %d", expected, result)
	}
}

func TestEvaluateComplexExpression(t *testing.T) {
	expr := "5 1 2 + 4 * + 3 -"
	expected := 14

	result, err := EvaluatePostfix(expr)
	if err != nil {
		t.Fatalf("Неочікувана помилка: %v", err)
	}
	if result != expected {
		t.Errorf("Очікувалось %d, отримано %d", expected, result)
	}
}

func TestEvaluateWithMultiplicationAndSubtraction(t *testing.T) {
	expr := "10 2 8 * + 3 -"
	expected := 23

	result, err := EvaluatePostfix(expr)
	if err != nil {
		t.Fatalf("Неочікувана помилка: %v", err)
	}
	if result != expected {
		t.Errorf("Очікувалось %d, отримано %d", expected, result)
	}
}

func TestEvaluateWithDivision(t *testing.T) {
	expr := "3 4 + 2 * 7 /"
	expected := 2

	result, err := EvaluatePostfix(expr)
	if err != nil {
		t.Fatalf("Неочікувана помилка: %v", err)
	}
	if result != expected {
		t.Errorf("Очікувалось %d, отримано %d", expected, result)
	}
}

func TestDivisionByZero(t *testing.T) {
	expr := "4 0 /"

	_, err := EvaluatePostfix(expr)
	if err == nil {
		t.Errorf("Очікувалась помилка при діленні на нуль")
	}
}

func TestInsufficientOperands(t *testing.T) {
	expr := "2 +"

	_, err := EvaluatePostfix(expr)
	if err == nil {
		t.Errorf("Очікувалась помилка через недостатню кількість операндів")
	}
}

func TestTooManyOperands(t *testing.T) {
	expr := "5 5"

	_, err := EvaluatePostfix(expr)
	if err == nil {
		t.Errorf("Очікувалась помилка через надлишок операндів")
	}
}

func TestInvalidToken(t *testing.T) {
	expr := "a b +"

	_, err := EvaluatePostfix(expr)
	if err == nil {
		t.Errorf("Очікувалась помилка через некоректні токени")
	}
}
