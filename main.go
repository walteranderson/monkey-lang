package main

import (
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	filename := os.Args[1]
	content, err := getFileContent(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	l := lexer.New(content)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Println("Parsing errors:")
		for _, e := range p.Errors() {
			fmt.Printf("\t%s\n", e)
		}
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}

func getFileContent(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
