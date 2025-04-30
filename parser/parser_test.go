package parser

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/Serein-sz/knife/lexer"
)

func TestParser(t *testing.T) {
	src, err := ReadFile("../example/parser.k")
	if err != nil {
		panic("not found source code")
	}
	l := lexer.New(src)
	p := New(l)
	_ = p.ParseProgram()
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
