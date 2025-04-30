package environment

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/Serein-sz/knife/ast"
)

type ObjectType string

type Object interface {
	Inspect() string
	Type() ObjectType
}

const (
	NUMBER          = "NUMBER"
	STRING          = "STRING"
	BOOLEAN         = "BOOLEAN"
	RETURN_VALUE    = "RETURN_VALUE"
	FUNCTION_DEFINE = "FUNCTION_DEFINE"
	BUILTIN         = "BUILTIN"
	NULL            = "NULL"
)

type HashKey struct {
	Type ObjectType
	Key  uint64
}

type Hashable interface {
	Object
	HashKey() HashKey
}

type Number struct {
	Value string
}

func (n *Number) Inspect() string {
	return fmt.Sprintf("%s", n.Value)
}

func (n *Number) Type() ObjectType {
	return NUMBER
}

func (n *Number) HashKey() HashKey {
	h := fnv.New64a()
	_, err := h.Write([]byte(n.Value))
	if err != nil {
		panic(err)
	}
	return HashKey{
		Type: n.Type(),
		Key:  h.Sum64(),
	}
}

type String struct {
	Value string
}

func (s *String) Inspect() string {
	return fmt.Sprintf("%s", s.Value)
}

func (s *String) Type() ObjectType {
	return STRING
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	_, err := h.Write([]byte(s.Value))
	if err != nil {
		panic(err)
	}
	return HashKey{
		Type: s.Type(),
		Key:  h.Sum64(),
	}
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%v", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN
}

type Null struct {
}

func (n *Null) Inspect() string {
	return "null"
}

func (b *Null) Type() ObjectType {
	return NULL
}

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Inspect() string {
	return r.Value.Inspect()
}

func (r *ReturnValue) Type() ObjectType {
	return RETURN_VALUE
}

type FunctionDefine struct {
	Parameters []*ast.Identifier
	Value      Object
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *FunctionDefine) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString("{")
	out.WriteString(f.Body.String())
	out.WriteString("}\n")

	return out.String()
}

func (f *FunctionDefine) Type() ObjectType {
	return FUNCTION_DEFINE
}

type Builtin struct {
	Name     string
	Function func(args ...Object) Object
}

func (b *Builtin) Inspect() string {
	return fmt.Sprintf("%v is a builtin\n", b.Name)
}

func (b *Builtin) Type() ObjectType {
	return BUILTIN
}
