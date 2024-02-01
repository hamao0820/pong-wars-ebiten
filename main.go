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

	numSquaresX = screenWidth / squareSize
	numSquaresY = screenHeight / squareSize

	dayColor   = color.RGBA{0xd9, 0xe8, 0xe3, 0xff}
	nightColor = color.RGBA{0x11, 0x4c, 0x5a, 0xff}
)

type Game struct {
	dayBall   *Ball
	nightBall *Ball

	squares [][]*Square
}

func newGame() *Game {
	dayBall := newBall(dayBallX, dayBallY, dayBallDx, dayBallDy, dayBallColor)
	nightBall := newBall(nightBallX, nightBallY, nightBallDx, nightBallDy, nightBallColor)

	squares := [][]*Square{}
	for i := 0; i < numSquaresX; i++ {
		squares = append(squares, []*Square{})
		for j := 0; j < numSquaresY; j++ {
			var color color.Color
			if i < numSquaresX/2 {
				color = dayColor
			} else {
				color = nightColor
			}
			square := newSquare(float32(i)*squareSize, float32(j)*squareSize, color)
			squares[i] = append(squares[i], square)
		}
	}

	return &Game{
		dayBall:   dayBall,
		nightBall: nightBall,
		squares:   squares,
	}
}

func (g *Game) Update() error {
	g.dayBall.Update()
	g.nightBall.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, row := range g.squares {
		for _, square := range row {
			square.Draw(screen)
		}
	}

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
