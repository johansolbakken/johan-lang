package generator

import (
	"fmt"
	"johan-lang/internal/parser"
	symboltable "johan-lang/internal/symbolTable"
	"os"
)

type PythonGenerator struct {
}

func NewPythonGenerator() *PythonGenerator {
	return &PythonGenerator{}
}

func (g *PythonGenerator) Generate(ast *parser.Node, stringList *symboltable.StringList) {
	filename := "main.py"

	// open file for writing
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	g.WriteStrings(file, stringList)

	g.WriteAst(file, ast)
}

func (g *PythonGenerator) WriteStrings(file *os.File, stringList *symboltable.StringList) {
	file.WriteString("# Strings\n")
	for i, item := range stringList.Items {
		str := fmt.Sprintf("str%d = \"%s\"\n", i, item)
		file.WriteString(str)
	}
}

func (g *PythonGenerator) WriteAst(file *os.File, node *parser.Node) {
	file.WriteString("# AST\n")
	g.WriteNode(file, node)
	file.WriteString("\n")
}

func (g *PythonGenerator) WriteNode(file *os.File, node *parser.Node) {
	if node.Type == parser.NodeType_PrintStatement {
		str := fmt.Sprintf("print(str%d)\n", node.Children[0].StringIndex)
		file.WriteString(str)
	}

	for _, child := range node.Children {
		g.WriteNode(file, child)
	}
}
