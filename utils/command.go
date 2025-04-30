package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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

// Format 格式化.k文件或递归格式化文件夹中的.k文件
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.1
func Format(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic("路径不存在")
	}

	if fileInfo.IsDir() {
		filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(filePath, ".k") {
				formatFile(filePath)
			}
			return nil
		})
	} else {
		formatFile(path)
	}
}

func formatFile(filePath string) {
	src, err := ReadFile(filePath)
	if err != nil {
		panic("未找到源代码文件")
	}
	l := lexer.New(src)
	p := parser.New(l)
	program := p.ParseProgram()
	if err = p.Error(); err != nil {
		io.WriteString(os.Stderr, err.Error())
	}
	err = WriteFile(filePath, program.String())
	if err != nil {
		io.WriteString(os.Stderr, err.Error())
	}
}
