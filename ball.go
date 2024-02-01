package main

import (
	"image/color"
	"math"

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

func (b *Ball) Update(squares [][]*Square) {
	dx := b.dx
	dy := b.dy

	for angle := 0.0; angle < math.Pi*2; angle += math.Pi / 4 {
		x := b.x + ballR*float32(math.Cos(angle))
		y := b.y + ballR*float32(math.Sin(angle))

		i := int(x / squareSize)
		j := int(y / squareSize)

		if i < 0 || i >= numSquaresX || j < 0 || j >= numSquaresY {
			continue
		}

		square := squares[i][j]
		if square.color != b.color {
			continue
		}

		if square.color == dayColor {
			square.color = nightColor
		} else {
			square.color = dayColor
		}

		if math.Abs(math.Cos(angle)) > math.Abs(math.Sin(angle)) {
			dx *= -1
		} else {
			dy *= -1
		}
	}

	b.dx = dx
	b.dy = dy

	if b.x+b.dx > screenWidth-ballR || b.x+b.dx < ballR {
		b.reflectX()
	}
	if b.y+b.dy > screenHeight-ballR || b.y+b.dy < ballR {
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
