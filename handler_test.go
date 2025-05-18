package postfixcli

import (
	"bytes"
	"strings"
	"testing"
)

func TestComputeHandler_Addition(t *testing.T) {
	testComputeValid(t, "3 4 +\n", "7\n")
}

func TestComputeHandler_Subtraction(t *testing.T) {
	testComputeValid(t, "10 4 -\n", "6\n")
}

func TestComputeHandler_Multiplication(t *testing.T) {
	testComputeValid(t, "2 3 *\n", "6\n")
}

func TestComputeHandler_Division(t *testing.T) {
	testComputeValid(t, "8 2 /\n", "4\n")
}

func TestComputeHandler_ComplexExpression(t *testing.T) {
	testComputeValid(t, "5 1 2 + 4 * + 3 -\n", "14\n")
}

func TestComputeHandler_DivisionByZero(t *testing.T) {
	testComputeInvalid(t, "4 0 /\n")
}

func TestComputeHandler_InvalidToken(t *testing.T) {
	testComputeInvalid(t, "3 a +\n")
}

func TestComputeHandler_NotEnoughOperands(t *testing.T) {
	testComputeInvalid(t, "4 +\n")
}

func TestComputeHandler_TooManyOperands(t *testing.T) {
	testComputeInvalid(t, "3 4 5 +\n")
}

func testComputeValid(t *testing.T, input, expected string) {
	in := strings.NewReader(input)
	out := &bytes.Buffer{}

	handler := ComputeHandler{Input: in, Output: out}
	err := handler.Compute()

	if err != nil {
		t.Fatalf("очікувався успішний результат, але отримано помилку: %v", err)
	}

	if out.String() != expected {
		t.Fatalf("очікувано %q, але отримано %q", expected, out.String())
	}
}

func testComputeInvalid(t *testing.T, input string) {
	in := strings.NewReader(input)
	out := &bytes.Buffer{}

	handler := ComputeHandler{Input: in, Output: out}
	err := handler.Compute()

	if err == nil {
		t.Fatalf("очікувалась помилка, але її не було")
	}
}
