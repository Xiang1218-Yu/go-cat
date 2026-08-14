package main

import (
	"image/color"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// Pet 桌面宠物结构
type Pet struct {
	window    fyne.Window
	image     *canvas.Image
	circle    *canvas.Circle
	x, y      float32
	vx, vy    float32
	edgeCount int
}

func main() {
	a := app.New()

	// 创建无边框窗口
	w := a.NewWindow("Desktop Pet")
	w.SetPadded(false)
	w.Resize(fyne.NewSize(100, 100))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	// 创建宠物实例
	pet := &Pet{
		window: w,
		x:      float32(rand.Intn(800)),
		y:      float32(rand.Intn(600)),
		vx:     2.0,
		vy:     1.5,
	}

	// 尝试加载图片资源
	imagePath := "assets/pet.png"
	if _, err := os.Stat(imagePath); err == nil {
		// 如果图片存在，使用图片
		pet.image = canvas.NewImageFromFile(imagePath)
		pet.image.FillMode = canvas.ImageFillOriginal
		pet.image.SetMinSize(fyne.NewSize(80, 80))
	} else {
		// 否则使用圆形作为默认宠物
		pet.circle = canvas.NewCircle(color.RGBA{R: 255, G: 182, B: 193, A: 255}) // 粉色
		pet.circle.StrokeColor = color.RGBA{R: 255, G: 105, B: 180, A: 255}         // 深粉色边框
		pet.circle.StrokeWidth = 3.0
		pet.circle.Resize(fyne.NewSize(80, 80))
	}

	// 创建容器
	container := container.NewWithoutLayout()
	if pet.image != nil {
		container.Add(pet.image)
	} else {
		container.Add(pet.circle)
	}

	w.SetContent(container)
	w.SetMaster()

	// 启动动画循环
	go pet.animate()

	// 启动随机移动
	go pet.randomMove()

	w.ShowAndRun()
}

// animate 处理宠物的动画效果
func (p *Pet) animate() {
	ticker := time.NewTicker(16 * time.Millisecond)
	for range ticker.C {
		// 更新位置
		p.x += p.vx
		p.y += p.vy

		// 获取屏幕尺寸（使用默认值）
		screenW := float32(1920)
		screenH := float32(1080)

		// 边界检测和反弹
		if p.x <= 0 || p.x >= screenW-100 {
			p.vx = -p.vx
			p.edgeCount++
		}
		if p.y <= 0 || p.y >= screenH-100 {
			p.vy = -p.vy
			p.edgeCount++
		}

		// 限制位置在屏幕范围内
		if p.x < 0 {
			p.x = 0
		}
		if p.x > screenW-100 {
			p.x = screenW - 100
		}
		if p.y < 0 {
			p.y = 0
		}
		if p.y > screenH-100 {
			p.y = screenH - 100
		}

		// 移动窗口
		p.window.Canvas().Content().Refresh()
	}
}

// randomMove 随机改变移动方向
func (p *Pet) randomMove() {
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		// 随机改变速度方向
		if rand.Float32() > 0.5 {
			p.vx = -p.vx
		}
		if rand.Float32() > 0.5 {
			p.vy = -p.vy
		}

		// 随机调整速度大小
		speed := 1.0 + rand.Float32()*2.0
		if p.vx > 0 {
			p.vx = speed
		} else {
			p.vx = -speed
		}
		if p.vy > 0 {
			p.vy = speed * 0.75
		} else {
			p.vy = -speed * 0.75
		}
	}
}