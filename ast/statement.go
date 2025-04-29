package ast

import (
	"bytes"

	"github.com/Serein-sz/knife/token"
)

// Statement 表示AST中的语句节点
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type Statement interface {
	Node
	statementNode()
}

// LetStatement 表示let语句节点
// 包含Token、变量名和表达式值
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral 返回let语句的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String 返回let语句的字符串表示
// 格式为: let <变量名> = <表达式>;
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.Token.Literal + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";\n")
	return out.String()
}

// statementNode 标记节点类型为语句
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (ls *LetStatement) statementNode() {}

// FunctionDefineStatement 表示函数定义语句节点
// 包含Token、函数名、参数列表和函数体
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type FunctionDefineStatement struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

// TokenLiteral 返回函数定义语句的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (fds *FunctionDefineStatement) TokenLiteral() string {
	return fds.Token.Literal
}

// String 返回函数定义语句的字符串表示
// 格式为: fn <函数名>(<参数列表>) {函数体}
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
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

// statementNode 标记节点类型为语句
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (fds *FunctionDefineStatement) statementNode() {}

// BlockStatement 表示代码块语句节点
// 包含Token和语句列表
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

// TokenLiteral 返回代码块语句的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

// String 返回代码块语句的字符串表示
//
//	格式为: {
//	    <语句1>
//	    <语句2>
//	}
//
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	for _, s := range bs.Statements {
		out.WriteString("    " + s.String() + "\n")
	}
	out.WriteString("}\n")
	return out.String()
}

// statementNode 标记节点类型为语句
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (fds *BlockStatement) statementNode() {}

// ReturnStatement 表示return语句节点
// 包含Token和返回值表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type ReturnStatement struct {
	Token token.Token
	Value Expression
}

// TokenLiteral 返回return语句的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String 返回return语句的字符串表示
// 格式为: return <表达式>;
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.Token.Literal + " ")
	out.WriteString(rs.Value.String() + ";")
	return out.String()
}

// statementNode 标记节点类型为语句
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (rs *ReturnStatement) statementNode() {}

// ExpressionStatement 表示表达式语句节点
// 包含Token和表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// TokenLiteral 返回表达式语句的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String 返回表达式语句的字符串表示
// 格式为: <表达式>
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (es *ExpressionStatement) String() string {
	return es.Expression.String()
}

// statementNode 标记节点类型为语句
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (es *ExpressionStatement) statementNode() {}
