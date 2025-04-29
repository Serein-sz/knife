# 支持的平台列表
$platforms = @(
  "windows/amd64",
  "linux/amd64",
  "darwin/amd64",
  "darwin/arm64"
)

# 遍历所有平台
foreach ($platform in $platforms) {
  # 分割 GOOS 和 GOARCH
  $parts = $platform -split "/"
  $GOOS = $parts[0]
  $GOARCH = $parts[1]

  # 生成输出文件名
  $output = "bin/app/app-$GOOS-$GOARCH"
  if ($GOOS -eq "windows") {
    $output += ".exe"
  }

  # 创建输出目录（如果不存在）
  if (-not (Test-Path "bin")) {
    New-Item -ItemType Directory -Path "bin" | Out-Null
  }

  # 执行编译
  Write-Host "正在编译 [$GOOS/$GOARCH]..."
  $env:GOOS = $GOOS
  $env:GOARCH = $GOARCH
  go build -o $output main.go
}

Write-Host "编译完成！"
