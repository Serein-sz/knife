package eval

import (
	"fmt"
	"strconv"
	"strings"
)

// AddNumberStrings 处理两个字符串数字相加，支持整数和浮点数
// 参数: num1, num2 - 要相加的数字字符串
// 返回: 相加结果的字符串表示和可能的错误
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.0
func AddNumberStrings(num1, num2 string) (string, error) {
	return CalculateNumbers(num1, num2, "+")
}

// SubtractNumberStrings 处理两个字符串数字相减，支持整数和浮点数
// 参数: num1, num2 - 要相减的数字字符串
// 返回: 相减结果的字符串表示和可能的错误
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.1
func SubtractNumberStrings(num1, num2 string) (string, error) {
	return CalculateNumbers(num1, num2, "-")
}

// MultiplyNumberStrings 处理两个字符串数字相乘，支持整数和浮点数
// 参数: num1, num2 - 要相乘的数字字符串
// 返回: 相乘结果的字符串表示和可能的错误
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.1
func MultiplyNumberStrings(num1, num2 string) (string, error) {
	return CalculateNumbers(num1, num2, "*")
}

// DivideNumberStrings 处理两个字符串数字相除，支持整数和浮点数
// 参数: num1, num2 - 要相除的数字字符串
// 返回: 相除结果的字符串表示和可能的错误
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.1
func DivideNumberStrings(num1, num2 string) (string, error) {
	return CalculateNumbers(num1, num2, "/")
}

// CalculateNumbers 处理两个字符串数字的四则运算
// 参数: num1, num2 - 要运算的数字字符串，op - 运算符(+, -, *, /)
// 返回: 运算结果的字符串表示和可能的错误
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.1
func CalculateNumbers(num1, num2, op string) (string, error) {
	// 检查是否为浮点数
	isFloat1 := strings.Contains(num1, ".")
	isFloat2 := strings.Contains(num2, ".")

	if isFloat1 || isFloat2 {
		// 处理浮点数运算
		f1, err := strconv.ParseFloat(num1, 64)
		if err != nil {
			return "", fmt.Errorf("无法解析第一个数字: %v", err)
		}

		f2, err := strconv.ParseFloat(num2, 64)
		if err != nil {
			return "", fmt.Errorf("无法解析第二个数字: %v", err)
		}

		var result float64
		switch op {
		case "+":
			result = f1 + f2
		case "-":
			result = f1 - f2
		case "*":
			result = f1 * f2
		case "/":
			if f2 == 0 {
				return "", fmt.Errorf("除数不能为零")
			}
			result = f1 / f2
		default:
			return "", fmt.Errorf("不支持的运算符: %s", op)
		}
		return strconv.FormatFloat(result, 'f', -1, 64), nil
	}

	// 处理整数运算
	i1, err := strconv.Atoi(num1)
	if err != nil {
		return "", fmt.Errorf("无法解析第一个数字: %v", err)
	}

	i2, err := strconv.Atoi(num2)
	if err != nil {
		return "", fmt.Errorf("无法解析第二个数字: %v", err)
	}

	var result int
	switch op {
	case "+":
		result = i1 + i2
	case "-":
		result = i1 - i2
	case "*":
		result = i1 * i2
	case "/":
		if i2 == 0 {
			return "", fmt.Errorf("除数不能为零")
		}
		result = i1 / i2
	default:
		return "", fmt.Errorf("不支持的运算符: %s", op)
	}
	return strconv.Itoa(result), nil
}
