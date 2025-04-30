package utils

import (
	"errors"
	"os"
	"strings"
)

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

func WriteFile(filename string, content string) error {
	// 检查文件扩展名是否为.k
	if !strings.HasSuffix(filename, ".k") {
		return errors.New("The file extension must be .k")
	}

	// 写入文件内容
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}
