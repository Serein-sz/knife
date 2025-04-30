package main

import (
	"flag"
	"fmt"

	"github.com/Serein-sz/knife/utils"
)

// main 程序入口
// 作者: 王强
// 日期: 2025-04-30
// 版本: 1.0.1
func main() {
	var (
		runPath         string
		runPathShort    string
		formatPath      string
		formatPathShort string
	)
	flag.StringVar(&runPath, "run", "", "主程序入口文件路径")
	flag.StringVar(&runPathShort, "r", "", "主程序入口文件路径(缩写)")
	flag.StringVar(&formatPath, "format", "", "需要格式化的程序文件路径")
	flag.StringVar(&formatPathShort, "f", "", "需要格式化的程序文件路径(缩写)")
	flag.Parse()

	if runPath != "" {
		utils.Run(runPath)
	} else if runPathShort != "" {
		utils.Run(runPathShort)
	} else if formatPath != "" {
		utils.Format(formatPath)
	} else if formatPathShort != "" {
		utils.Format(formatPathShort)
	} else {
		fmt.Println("请使用-run/-r或-format/-f参数指定操作")
		flag.Usage()
	}
}
