package main

import (
	"math/rand"
)

type Board struct {
	data       []bool
	swapBuffer []bool
	width      int
	height     int
}

func NewBoard(width, height int) Board {

	data := make([]bool, width*height)
	for index := range data {
		data[index] = rand.Intn(2) == 0
	}

	return Board{
		data:       data,
		swapBuffer: make([]bool, width*height),
		width:      width,
		height:     height,
	}
}

func (board *Board) Update() {

	clear(board.swapBuffer)

	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			count := board.getNeighbourCount(x, y)
			index := y*board.width + x

			switch {
			case count < 2 || count > 3: // dies
				board.swapBuffer[index] = false
			case count == 3 || board.data[index]: // lives
				board.swapBuffer[index] = true
			}
		}
	}

	board.data, board.swapBuffer = board.swapBuffer, board.data
}

func (board *Board) getNeighbourCount(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {

			if dx == 0 && dy == 0 {
				continue // skip if middle pixel
			}

			index := (y+dy)*board.width + (x + dx)

			if index < 0 || index >= len(board.data) {
				continue // skip index if outside of domain
			}

			if board.data[index] {
				count++
			}
		}
	}

	return count
}

func (board *Board) ToPixels(pixels []byte) {
	for index, value := range board.data {

		color := byte(0x00)
		if value {
			color = 0xff
		}

		pixels[4*index+0] = color
		pixels[4*index+1] = color
		pixels[4*index+2] = color
		pixels[4*index+3] = 0xff
	}
}
