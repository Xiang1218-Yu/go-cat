#!/bin/bash

# 桌面宠物编译脚本

echo "开始编译桌面宠物..."

# 编译Windows版本
echo "编译Windows版本..."
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -s -w" -o desktop-pet-windows.exe main.go

# 编译Mac版本
echo "编译Mac版本..."
GOOS=darwin GOARCH=amd64 go build -o desktop-pet-macos-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -o desktop-pet-macos-arm64 main.go

echo "编译完成！"
echo "Windows版本: desktop-pet-windows.exe"
echo "Mac版本: desktop-pet-macos-amd64 (Intel)"
echo "Mac版本: desktop-pet-macos-arm64 (Apple Silicon)"