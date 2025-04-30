package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/Serein-sz/knife/environment"
	"github.com/Serein-sz/knife/eval"
	"github.com/Serein-sz/knife/lexer"
	"github.com/Serein-sz/knife/parser"
)

func Run(mainProgramPath string) {
	src, err := ReadFile(mainProgramPath)
	if err != nil {
		panic("not found source code")
	}
	l := lexer.New(src)
	p := parser.New(l)
	program := p.ParseProgram()
	if err = p.Error(); err != nil {
		io.WriteString(os.Stderr, err.Error())
	}
	env := environment.NewEnvironment(nil)
	_, err = eval.Eval(program, env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "eval err: %v", err)
	}
}

func Format(mainProgramPath string) {
	src, err := ReadFile(mainProgramPath)
	if err != nil {
		panic("not found source code")
	}
	l := lexer.New(src)
	p := parser.New(l)
	program := p.ParseProgram()
	if err = p.Error(); err != nil {
		io.WriteString(os.Stderr, err.Error())
	}
	err = WriteFile(mainProgramPath, program.String())
	if err != nil {
		io.WriteString(os.Stderr, err.Error())
	}
}
