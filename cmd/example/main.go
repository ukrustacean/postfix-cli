package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	postfixcli "github.com/ukrustacean/postfix-cli"
)

func main() {
	exprFlag := flag.String("e", "", "Постфіксний вираз для обчислення")
	fileFlag := flag.String("f", "", "Файл з вхідним виразом")
	outFlag := flag.String("o", "", "Файл для запису результату")

	flag.Parse()

	if *exprFlag != "" && *fileFlag != "" {
		log.Fatalln("Помилка: не можна використовувати одночасно -e та -f")
	}

	var input io.Reader
	var output io.Writer = os.Stdout

	if *exprFlag != "" {
		input = strings.NewReader(*exprFlag + "\n")
	} else if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			log.Fatalf("Не вдалося відкрити файл %s: %v\n", *fileFlag, err)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	if *outFlag != "" {
		file, err := os.Create(*outFlag)
		if err != nil {
			log.Fatalf("Не вдалося створити файл %s: %v\n", *outFlag, err)
		}
		defer file.Close()
		output = file
	}

	handler := &postfixcli.ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Помилка обчислення:", err)
		os.Exit(1)
	}
}
