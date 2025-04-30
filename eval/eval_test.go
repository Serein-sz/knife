package eval

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Serein-sz/knife/environment"
	"github.com/Serein-sz/knife/lexer"
	"github.com/Serein-sz/knife/parser"
)

func TestEval(t *testing.T) {
	src, err := ReadFile("../example/main.k")
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

func ReadFile(filename string) (string, error) {
	// 检查文件扩展名是否为.k
	if !strings.HasSuffix(filename, ".k") {
		return "", errors.New("The file extension must be .k")
	}

	// 读取文件内容
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// 确保使用UTF-8编码和LF换行符
	return string(content), nil
}
