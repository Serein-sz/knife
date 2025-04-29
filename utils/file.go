package utils

import (
	"errors"
	"os"
	"strings"
)

// readFile 读取指定路径的.k文件内容
// 作者: 王强
// 日期: 2025-04-29
// 版本: 1.0.0
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
