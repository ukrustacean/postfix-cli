package main

import (
	"flag"
	"log"
	"os"

	postfixcli "github.com/ukrustacean/postfix-cli"
)

var (
// inputExpression = flag.String("e", "", "Expression to compute")
// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	handler := &postfixcli.ComputeHandler{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
	err := handler.Compute()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
