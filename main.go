package main

import (
	"fmt"
	"johan-lang/internal/fileio"
	"johan-lang/internal/generator"
	"johan-lang/internal/lexer"
	"johan-lang/internal/parser"
	symboltable "johan-lang/internal/symbolTable"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	source := fileio.ReadWholeFile(filename)

	lex := lexer.NewLexer(source)
	tokens := lex.Lex()

	pars := parser.NewParser(tokens)
	ast := pars.Parse()

	stringList := symboltable.NewStringList()
	symboltable.ExtractStringFromAst(stringList, &ast)

	generator := generator.NewPythonGenerator()
	generator.Generate(&ast, stringList)
}
