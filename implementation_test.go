package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluatePostfix(t *testing.T) {
	res, err := EvaluatePostfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, 11, res)
	}
}

func ExampleEvaluatePostfix() {
	res, _ := EvaluatePostfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 4
}
