package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Square struct {
	x, y  float32
	color color.Color
}

func newSquare(x, y float32, color color.Color) *Square {
	return &Square{x: x, y: y, color: color}
}

func (s *Square) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, s.x, s.y, squareSize, squareSize, s.color, true)
}
