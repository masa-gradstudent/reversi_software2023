package main

import (
	"fmt"
)

const (
	WHITE = 1
	BLACK = -1
)

type Board struct {
	tokens      [8][8]int
	currentTurn int
}

func create_board() Board {
	b := Board{}
	b.tokens = [8][8]int{}
	return b
}

func defalut_board() Board {
	b := create_board()
	b.tokens[3][3] = WHITE
	b.tokens[3][4] = BLACK
	b.tokens[4][3] = BLACK
	b.tokens[4][4] = WHITE
	return b
}

func (b *Board) count() (int, int) {
	white := 0
	black := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.tokens[i][j] == WHITE {
				white += 1
			} else if b.tokens[i][j] == BLACK {
				black += 1
			}
		}
	}
	return white, black
}

func (b *Board) flip(row int, column int) {
	b.flip_forward(row, column, 1)
	b.flip_forward(row, column, -1)
}

func (b *Board) flip_forward(row int, column int, forward int) {
	color := b.tokens[row][column]
	flip_list := make([]int, 0, 8)
	for {
		row += forward
		if 0 > row || row >= 8 {
			break
		}
		ocolor := b.tokens[row][column]
		if -ocolor == color {
			flip_list = append(flip_list, row)
		} else {
			break
		}
	}
	for _, i := range flip_list {
		b.tokens[i][column] = color
	}
}

func determineWinner(white, black int) string {
	if white > black {
		return "White wins!"
	} else if black > white {
		return "Black wins!"
	} else {
		return "It's a tie!"
	}
}
func printBoard(b Board) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.tokens[i][j] == WHITE {
				fmt.Print("W ")
			} else if b.tokens[i][j] == BLACK {
				fmt.Print("B ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) switchTurn() {
	if b.currentTurn == WHITE {
		b.currentTurn = BLACK
	} else {
		b.currentTurn = WHITE
	}
}
func (b *Board) isValidMove(row int, column int) bool {
	if b.tokens[row][column] != 0 {
		return false
	}
	color := b.currentTurn
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			if b.checkDirection(row, column, dr, dc, color) {
				return true
			}
		}
	}
	return false
}

func (b *Board) checkDirection(row, column, dr, dc, color int) bool {
	row += dr
	column += dc
	if row < 0 || row >= 8 || column < 0 || column >= 8 {
		return false
	}
	ocolor := b.tokens[row][column]
	if ocolor == -color {
		return b.checkDirection(row, column, dr, dc, color)
	} else if ocolor == color {
		return true
	}
	return false
}

func (b *Board) makeMove(row, column int) bool {
	if !b.isValidMove(row, column) {
		return false
	}
	b.tokens[row][column] = b.currentTurn
	b.flip(row, column)
	b.currentTurn = -b.currentTurn
	return true
}
