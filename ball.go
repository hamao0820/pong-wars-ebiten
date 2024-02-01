package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	x, y, dx, dy float32
	color        color.Color
}

func newBall(x, y, dx, dy float32, color color.Color) *Ball {
	return &Ball{x: x, y: y, dx: dx, dy: dy, color: color}
}

func (b *Ball) Update() {
	if b.x < ballR || b.x > screenWidth-ballR {
		b.reflectX()
	}
	if b.y < ballR || b.y > screenHeight-ballR {
		b.reflectY()
	}

	b.x += b.dx
	b.y += b.dy
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.x, b.y, ballR, b.color, true)
}

func (b *Ball) reflectX() {
	b.dx *= -1
}

func (b *Ball) reflectY() {
	b.dy *= -1
}
