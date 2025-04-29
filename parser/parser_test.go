package parser

import (
	"fmt"
	"testing"

	"github.com/Serein-sz/knife/lexer"
	"github.com/Serein-sz/knife/utils"
)

func TestParser(t *testing.T)  {
	src, err := utils.ReadFile("../example/main.k")
	if err != nil {
		panic("not found source code")
	}
	l := lexer.New(src)
	p := New(l)
	program := p.ParseProgram()
	fmt.Printf("src:\n%+v", program)
}
