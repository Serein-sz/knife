#!/bin/bash

# 支持的平台列表
platforms=(
  "windows/amd64"
  "linux/amd64"
  "darwin/amd64"
  "darwin/arm64"  # M1/M2 芯片的 Mac
)

for platform in "${platforms[@]}"; do
  # 分割平台字符串为 GOOS 和 GOARCH
  GOOS=${platform%/*}
  GOARCH=${platform#*/}
  
  # 生成输出文件名
  output_name="./bin/app/app-${GOOS}-${GOARCH}"
  if [ $GOOS = "windows" ]; then
    output_name+=".exe"
  fi

  # 执行编译
  echo "Building for $GOOS/$GOARCH..."
  GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name main.go
done
