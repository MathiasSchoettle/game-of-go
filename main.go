package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const ScreenWidth = 1280
const ScreenHeight = 720

const GameWidth = 320
const GameHeight = 180

type Game struct {
	board  Board
	width  int
	height int
	pixels []byte
}

func (g *Game) Update() error {
	g.board.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.ToPixels(g.pixels)
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return GameWidth, GameHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Game of Life")

	board := NewBoard(GameWidth, GameHeight)
	game := Game{
		board:  board,
		width:  GameWidth,
		height: GameHeight,
		pixels: make([]byte, 4*GameWidth*GameHeight),
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
