package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 800

	squareSize = 25
)

var (
	dayBallX     = float32(screenWidth) / 4
	dayBallY     = float32(screenHeight) / 2
	dayBallDx    = float32(12.5)
	dayBallDy    = float32(-12.5)
	dayBallColor = color.RGBA{0x11, 0x4c, 0x5a, 0xff}

	nightBallX     = float32(screenWidth) * 3 / 4
	nightBallY     = float32(screenHeight) / 2
	nightBallDx    = float32(-12.5)
	nightBallDy    = float32(12.5)
	nightBallColor = color.RGBA{0xd9, 0xe8, 0xe3, 0xff} //#D9E8E3

	ballR = float32(squareSize) / 2
)

type Game struct {
	dayBall   *Ball
	nightBall *Ball
}

func newGame() *Game {
	g := &Game{}
	g.dayBall = newBall(dayBallX, dayBallY, dayBallDx, dayBallDy, dayBallColor)
	g.nightBall = newBall(nightBallX, nightBallY, nightBallDx, nightBallDy, nightBallColor)
	return g
}

func (g *Game) Update() error {
	g.dayBall.Update()
	g.nightBall.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.dayBall.Draw(screen)
	g.nightBall.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := newGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong Wars")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
