package ast

import (
	"bytes"

	"github.com/Serein-sz/knife/token"
)

// Expression 表示AST中的表达式节点接口
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type Expression interface {
	Node
	expressionNode()
}

// Identifier 表示标识符节点
// 包含Token和标识符值
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral 返回标识符的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String 返回标识符的字符串表示
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (i *Identifier) String() string {
	return i.Value
}

// expressionNode 标记节点类型为表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (i *Identifier) expressionNode() {}

// Null 表示空值节点
// 包含Token和标识符值
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type Null struct {
	Token token.Token
	Value string
}

// TokenLiteral 返回标识符的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (n *Null) TokenLiteral() string {
	return n.Token.Literal
}

// String 返回标识符的字符串表示
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (n *Null) String() string {
	return n.Value
}

// expressionNode 标记节点类型为表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (n *Null) expressionNode() {}

// NumberLiteral 表示数字字面量节点
// 包含Token和数字值
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type NumberLiteral struct {
	Token token.Token
	Value string
}

// TokenLiteral 返回数字字面量的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (nl *NumberLiteral) TokenLiteral() string {
	return nl.Token.Literal
}

// String 返回数字字面量的字符串表示
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (nl *NumberLiteral) String() string {
	return nl.Value
}

// expressionNode 标记节点类型为表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (nl *NumberLiteral) expressionNode() {}

// NumberLiteral 表示数字字面量节点
// 包含Token和数字值
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type StringLiteral struct {
	Token token.Token
	Value string
}

// TokenLiteral 返回字符串字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

// String 返回字符串字面量的字符串表示
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (sl *StringLiteral) String() string {
	return sl.Value
}

// expressionNode 标记节点类型为表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (sl *StringLiteral) expressionNode() {}

// FunctionCallExpression 表示函数调用表达式节点
// 包含Token、参数列表和函数表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type FunctionCallExpression struct {
	Token     token.Token
	Arguments []Expression
	Function  Expression
}

// TokenLiteral 返回函数调用表达式的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (ce *FunctionCallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String 返回函数调用表达式的字符串表示
// 格式为: <函数名>(<参数1>, <参数2>)
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
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

// expressionNode 标记节点类型为表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (ce *FunctionCallExpression) expressionNode() {}

// PrefixExpression 表示前缀表达式节点
// 包含Token、操作符和右表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type PrefixExpression struct {
	Token token.Token
	Op    string
	Rhs   Expression
}

// TokenLiteral 返回前缀表达式的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String 返回前缀表达式的字符串表示
// 格式为: <操作符><右表达式>
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString(pe.Op)
	out.WriteString(pe.Rhs.String())
	return out.String()
}

// expressionNode 标记节点类型为表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (ie *PrefixExpression) expressionNode() {}

// InfixExpression 表示中缀表达式节点
// 包含Token、左表达式、操作符和右表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
type InfixExpression struct {
	Token token.Token
	Lhs   Expression
	Op    string
	Rhs   Expression
}

// TokenLiteral 返回中缀表达式的字面量
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String 返回中缀表达式的字符串表示
// 格式为: <左表达式> <操作符> <右表达式>
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ie.Lhs.String() + " ")
	out.WriteString(ie.Op + " ")
	out.WriteString(ie.Rhs.String())
	return out.String()
}

// expressionNode 标记节点类型为表达式
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
func (ie *InfixExpression) expressionNode() {}
