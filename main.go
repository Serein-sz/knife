package main

import "github.com/Serein-sz/knife/utils"

func main() {
	mainProgramPath := "./example/main.k"
	utils.Run(mainProgramPath) // 运行主程序
	utils.Format(mainProgramPath)
}
