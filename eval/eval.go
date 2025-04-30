package eval

import (
	"fmt"

	"github.com/Serein-sz/knife/ast"
	"github.com/Serein-sz/knife/environment"
)

var (
	NULL  = &environment.Null{}
	TRUE  = &environment.Boolean{Value: true}
	FALSE = &environment.Boolean{Value: false}
)

func Eval(node ast.Node, env *environment.Environment) (environment.Object, error) {
	switch node := (node).(type) {
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.BlockStatement:
		return evalBlockStatements(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.LetStatement:
		return evalLetStatement(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.NumberLiteral:
		return &environment.Number{Value: node.Value}, nil
	case *ast.FunctionDefineStatement:
		return evalFunctionDefineStatement(node, env)
	case *ast.ReturnStatement:
		value, err := Eval(node.Value, env)
		return &environment.ReturnValue{Value: value}, err
	case *ast.FunctionCallExpression:
		function, err := Eval(node.Function, env)
		if err != nil {
			return nil, err
		}
		args, err := evalExpressions(node.Arguments, env)
		if err != nil {
			return nil, err
		}
		return evalFunctionCallExpression(function, args)
	case *ast.InfixExpression:
		lhs, err := Eval(node.Lhs, env)
		if err != nil {
			return nil, err
		}
		rhs, err := Eval(node.Rhs, env)
		if err != nil {
			return nil, err
		}
		res, err := evalInfixExpression(node.Op, lhs, rhs)
		return res, err
	}
	return nil, fmt.Errorf("unsupported object type: %T\n", node)
}

func evalBlockStatements(statements []ast.Statement, env *environment.Environment) (environment.Object, error) {
	var res environment.Object
	var err error
	for _, s := range statements {
		res, err = Eval(s, env)
		if err != nil {
			return nil, err
		}
		if r, ok := res.(*environment.ReturnValue); ok {
			return r, err
		}
	}
	return res, nil
}

func evalInfixExpression(op string, lhs environment.Object, rhs environment.Object) (environment.Object, error) {
	lType, rType := lhs.Type(), rhs.Type()
	if lType == environment.NUMBER && rType == environment.NUMBER {
		l, r := lhs.(*environment.Number), rhs.(*environment.Number)
		return evalInfixNumber(op, l, r)
	}
	return nil, fmt.Errorf("illegal operands for %q, lhs: %q, rhs: %q\n", op, lhs.Inspect(), rhs.Inspect())
}

func evalInfixNumber(op string, l *environment.Number, r *environment.Number) (environment.Object, error) {
	switch op {
	case "+":
		number, err := AddNumberStrings(l.Value, r.Value)
		return &environment.Number{Value: number}, err
	case "-":
		number, err := SubtractNumberStrings(l.Value, r.Value)
		return &environment.Number{Value: number}, err
	case "*":
		number, err := MultiplyNumberStrings(l.Value, r.Value)
		return &environment.Number{Value: number}, err
	case "/":
		number, err := DivideNumberStrings(l.Value, r.Value)
		return &environment.Number{Value: number}, err
	case "==":
		return &environment.Boolean{Value: l.Value == r.Value}, nil
	case "!=":
		return &environment.Boolean{Value: l.Value != r.Value}, nil
	}
	return nil, fmt.Errorf("unsupported infix operator for strings: %q %s %q\n", l.Inspect(), op, r.Inspect())
}

func evalProgram(statements []ast.Statement, env *environment.Environment) (environment.Object, error) {
	var result environment.Object
	var err error
	for _, statement := range statements {
		result, err = Eval(statement, env)
		if err != nil {
			return nil, err
		}
		if r, ok := result.(*environment.ReturnValue); ok {
			return r.Value, err
		}
	}
	return result, nil
}

func evalIdentifier(node *ast.Identifier, env *environment.Environment) (environment.Object, error) {
	if obj, err := env.Get(node.Value); err == nil {
		return obj, nil
	}

	bti, ok := builtins[node.Value]
	if ok {
		return bti, nil
	}

	return nil, fmt.Errorf("undefined identifier: %s\n", node.Value)
}

func evalLetStatement(node *ast.LetStatement, env *environment.Environment) (environment.Object, error) {
	obj, err := Eval(node.Value, env)
	if err != nil {
		return nil, err
	}
	env.Set(node.Name.Value, obj)
	return nil, nil
}

func evalFunctionDefineStatement(node *ast.FunctionDefineStatement, env *environment.Environment) (environment.Object, error) {
	params := node.Parameters
	body := node.Body
	functionDefine := &environment.FunctionDefine{
		Parameters: params,
		Body:       body,
		Env:        env,
	}
	return env.Set(node.Name.Value, functionDefine)
}

func evalExpressions(args []ast.Expression, env *environment.Environment) ([]environment.Object, error) {
	var res = make([]environment.Object, 0, len(args))
	for _, a := range args {
		v, err := Eval(a, env)
		if err != nil {
			return nil, fmt.Errorf("passing exp error: [%v]%v", a, err)
		}
		res = append(res, v)
	}
	return res, nil
}

func evalFunctionCallExpression(function environment.Object, args []environment.Object) (environment.Object, error) {
	switch f := function.(type) {
	case *environment.FunctionDefine:
		newEnv := environment.NewEnvironment(f.Env)
		for i, p := range f.Parameters {
			newEnv.Set(p.Value, args[i])
		}

		val, err := Eval(f.Body, newEnv)
		if err != nil {
			return nil, err
		}
		if v, ok := val.(*environment.ReturnValue); ok {
			return v.Value, nil
		}
		return val, nil
	case *environment.Builtin:
		return f.Function(args...), nil
	}
	return nil, fmt.Errorf("%v is not callable", function.Inspect())
}
