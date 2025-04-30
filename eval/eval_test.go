package eval

import (
	"io"
	"os"
	"testing"

	"github.com/Serein-sz/knife/environment"
	"github.com/Serein-sz/knife/lexer"
	"github.com/Serein-sz/knife/parser"
	"github.com/Serein-sz/knife/utils"
)

func TestEval(t *testing.T) {
	src, err := utils.ReadFile("../example/main.k")
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
	_, err = Eval(program, env)
	if err != nil {
		t.Fatalf("eval err: %v", err)
	}
}
