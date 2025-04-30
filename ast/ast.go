package ast

import "bytes"

// Node 表示抽象语法树中的节点接口
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.0
type Node interface {
	TokenLiteral() string
	String() string
}

// Program 表示程序节点，包含语句集合
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.0
type Program struct {
	Statements []Statement
}

// String 返回程序的字符串表示
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.1
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
