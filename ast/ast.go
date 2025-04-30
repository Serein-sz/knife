package ast

import "bytes"

type Node interface {
	Line() int
	TokenLiteral() string
	String() string
}

type Program struct {
	Statements []Statement
}

func (p *Program) Line() int {
	return p.Statements[0].Line()
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	return p.Statements[0].TokenLiteral()
}
