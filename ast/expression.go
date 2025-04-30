package ast

import (
	"bytes"
	"strings"

	"github.com/Serein-sz/knife/token"
)

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) Line() int {
	return i.Token.Line
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}

type Null struct {
	Token token.Token
	Value string
}

func (n *Null) Line() int {
	return n.Token.Line
}

func (n *Null) TokenLiteral() string {
	return n.Token.Literal
}

func (n *Null) String() string {
	return n.Value
}

func (n *Null) expressionNode() {}

type NumberLiteral struct {
	Token token.Token
	Value string
}

func (n *NumberLiteral) Line() int {
	return n.Token.Line
}

func (nl *NumberLiteral) TokenLiteral() string {
	return nl.Token.Literal
}

func (nl *NumberLiteral) String() string {
	return nl.Value
}

func (nl *NumberLiteral) expressionNode() {}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (s *StringLiteral) Line() int {
	return s.Token.Line
}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) String() string {
	return sl.Value
}

func (sl *StringLiteral) expressionNode() {}

type FunctionCallExpression struct {
	Token     token.Token
	Arguments []Expression
	Function  Expression
}

func (fce *FunctionCallExpression) Line() int {
	return fce.Token.Line
}

func (ce *FunctionCallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *FunctionCallExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	for index, expression := range ce.Arguments {
		out.WriteString(strings.ReplaceAll(expression.String(), "\n", ""))
		if index != len(ce.Arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")\n")
	return out.String()
}

func (ce *FunctionCallExpression) expressionNode() {}

type PrefixExpression struct {
	Token token.Token
	Op    string
	Rhs   Expression
}

func (pe *PrefixExpression) Line() int {
	return pe.Token.Line
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString(pe.Op)
	out.WriteString(pe.Rhs.String())
	return out.String()
}

func (ie *PrefixExpression) expressionNode() {}

type InfixExpression struct {
	Token token.Token
	Lhs   Expression
	Op    string
	Rhs   Expression
}

func (ie *InfixExpression) Line() int {
	return ie.Token.Line
}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ie.Lhs.String() + " ")
	out.WriteString(ie.Op + " ")
	out.WriteString(ie.Rhs.String())
	return out.String()
}

func (ie *InfixExpression) expressionNode() {}
