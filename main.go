package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Serein-sz/knife/environment"
	"github.com/Serein-sz/knife/eval"
	"github.com/Serein-sz/knife/lexer"
	"github.com/Serein-sz/knife/parser"
	"github.com/Serein-sz/knife/utils"
)

func main() {
	src, err := utils.ReadFile("./example/main.k")
	if err != nil {
		panic("not found source code")
	}
	l := lexer.New(src)
	p := parser.New(l)
	program := p.ParseProgram()
	if err = p.Error(); err != nil {
		io.WriteString(os.Stderr, err.Error())
	}
	fmt.Println(program)
	env := environment.NewEnvironment(nil)
	_, err = eval.Eval(program, env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "eval err: %v", err)
	}
}
