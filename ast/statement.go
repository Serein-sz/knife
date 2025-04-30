package ast

import (
	"bytes"
	"strings"

	"github.com/Serein-sz/knife/token"
)

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) Line() int {
	return ls.Token.Line
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.Token.Literal + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString("\n")
	return out.String()
}

func (ls *LetStatement) statementNode() {}

type FunctionDefineStatement struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fds *FunctionDefineStatement) Line() int {
	return fds.Token.Line
}

func (fds *FunctionDefineStatement) TokenLiteral() string {
	return fds.Token.Literal
}

func (fds *FunctionDefineStatement) String() string {
	var out bytes.Buffer
	out.WriteString(fds.Token.Literal + " ")
	out.WriteString(fds.Name.String())
	out.WriteString("(")
	for index, identifier := range fds.Parameters {
		out.WriteString(identifier.String())
		if index != len(fds.Parameters)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(") ")
	out.WriteString(fds.Body.String())
	return out.String()
}

func (fds *FunctionDefineStatement) statementNode() {}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) Line() int {
	return bs.Token.Line
}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	for i, s := range bs.Statements {
		out.WriteString("    " + strings.ReplaceAll(s.String(), "\n", ""))
		if i != len(bs.Statements)-1 {
			out.WriteString("\n")
		}
	}
	out.WriteString("\n}\n")
	return out.String()
}

func (fds *BlockStatement) statementNode() {}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) Line() int {
	return rs.Token.Line
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.Token.Literal + " ")
	out.WriteString(rs.Value.String())
	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) Line() int {
	return es.Token.Line
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	return es.Expression.String()
}

func (es *ExpressionStatement) statementNode() {}
