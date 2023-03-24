package ir

import "johan-lang/internal/parser"

type OpCode int
type Operand int

type Tac struct {
	OpCode   OpCode
	OperandA Operand
	OperandB Operand
	Result   Operand
}

const (
	OpCode_Add OpCode = iota
	OpCode_Sub
	OpCode_Mul
	OpCode_Div
	OpCode_Mod
	OpCode_Eq
	OpCode_Ne
	OpCode_Gt
	OpCode_Ge
	OpCode_Lt
	OpCode_Le
	OpCode_And
	OpCode_Or
	OpCode_Not
	OpCode_Print
)

const (
	Operand_Unknown Operand = iota
	Operand_Identifier
	Operand_String
	Operand_Number
	Operand_RegisterA
)

type TacGenerator struct {
	Instructions []Tac
}

func NewTacGenerator() *TacGenerator {
	return &TacGenerator{}
}

func (t *TacGenerator) Generate(node *parser.Node) {
	switch node.Type {
	case parser.NodeType_PrintStatement:
		t.GeneratePrintStatement(node)
	}
}

func (t *TacGenerator) GeneratePrintStatement(node *parser.Node) {
	t.Generate(node.Children[1])
	t.Instructions = append(t.Instructions, Tac{
		OpCode:   OpCode_Print,
		OperandA: Operand_Unknown,
		OperandB: Operand_Unknown,
		Result:   Operand_Unknown,
	})
}
