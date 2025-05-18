package postfixcli

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.Input)
	expr, err := reader.ReadString('\n')

	expr = strings.TrimSpace(expr)

	if err != nil {
		return err
	}

	result, err := EvaluatePostfix(expr)

	if err != nil {
		return err
	}

	s := strconv.Itoa(result)

	writer := bufio.NewWriter(ch.Output)
	_, err = writer.WriteString(s + "\n")

	if err != nil {
		return err
	}

	err = writer.Flush()

	if err != nil {
		return err
	}

	return nil
}
