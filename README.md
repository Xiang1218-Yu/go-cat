# 桌面宠物

一个使用Go语言和Fyne框架开发的跨平台桌面宠物应用。

## 功能特性

- 跨平台支持（Windows和Mac）
- 无边框窗口设计
- 自动移动和反弹动画
- 随机改变移动方向
- 支持自定义宠物图片

## 环境要求

- Go 1.21 或更高版本
- Fyne框架依赖

### Mac系统依赖

```bash
# 安装Xcode命令行工具
xcode-select --install

# 安装其他依赖
brew install pkg-config
```

### Windows系统依赖

需要安装MinGW-w64：
- 下载并安装MinGW-w64
- 确保gcc编译器在PATH中

## 安装依赖

```bash
go mod download
```

## 编译

### 使用编译脚本

```bash
chmod +x build.sh
./build.sh
```

### 手动编译

#### Windows版本
```bash
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -s -w" -o desktop-pet-windows.exe main.go
```

#### Mac版本
```bash
# Intel Mac
GOOS=darwin GOARCH=amd64 go build -o desktop-pet-macos-amd64 main.go

# Apple Silicon Mac
GOOS=darwin GOARCH=arm64 go build -o desktop-pet-macos-arm64 main.go
```

## 自定义宠物图片

1. 创建`assets`文件夹
2. 将宠物图片命名为`pet.png`放入`assets`文件夹
3. 重新编译程序

如果不提供图片，程序会使用默认的粉色圆形作为宠物。

## 运行

直接运行编译后的可执行文件：

```bash
# Windows
desktop-pet-windows.exe

# Mac
./desktop-pet-macos-amd64  # Intel Mac
./desktop-pet-macos-arm64  # Apple Silicon Mac
```

## 特性说明

- **无边框窗口**：宠物窗口没有标题栏和边框
- **自动移动**：宠物会在屏幕上自动移动
- **边界反弹**：碰到屏幕边缘会反弹
- **随机行为**：每3秒随机改变移动方向和速度
- **置顶显示**：宠物窗口始终显示在最前面

## 技术栈

- **语言**：Go 1.21
- **GUI框架**：Fyne v2.4.3
- **跨平台**：支持Windows、MacOS

## 许可证

MIT License