@echo off
echo 开始编译桌面宠物...

echo 编译Windows版本...
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-H windowsgui -s -w" -o desktop-pet-windows.exe main.go

echo 编译Mac版本...
set GOOS=darwin
set GOARCH=amd64
go build -o desktop-pet-macos-amd64 main.go

set GOARCH=arm64
go build -o desktop-pet-macos-arm64 main.go

echo 编译完成！
echo Windows版本: desktop-pet-windows.exe
echo Mac版本: desktop-pet-macos-amd64 (Intel)
echo Mac版本: desktop-pet-macos-arm64 (Apple Silicon)

pause