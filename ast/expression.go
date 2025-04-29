package ast

import (
	"bytes"

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

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}

type NumberLiteral struct {
	Token token.Token
	Value string
}

func (nl *NumberLiteral) TokenLiteral() string {
	return nl.Token.Literal
}

func (nl *NumberLiteral) String() string {
	return nl.Value
}

func (nl *NumberLiteral) expressionNode() {}

type FunctionCallExpression struct {
	Token     token.Token
	Arguments []Expression
	Function  Expression
}

func (ce *FunctionCallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *FunctionCallExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	for index, expression := range ce.Arguments {
		out.WriteString(expression.String())
		if index != len(ce.Arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")
	return out.String()
}

func (ce *FunctionCallExpression) expressionNode() {}

type PrefixExpression struct {
	Token token.Token
	Op    string
	Rhs   Expression
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
