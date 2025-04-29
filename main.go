package main

import (
	"fmt"
	
	"github.com/Serein-sz/knife/lexer"
	"github.com/Serein-sz/knife/parser"
	"github.com/Serein-sz/knife/utils"
)

func main() {
	src, err := utils.ReadFile("D:/go-workspace/knife/example/main.k")
	if err != nil {
		panic("not found source code")
	}
	l := lexer.New(src)
	p := parser.New(l)
	program := p.ParseProgram()
	fmt.Printf("src:\n%+v", program)
}
