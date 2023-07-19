package main

const (
	WHITE = 1
	BLACK = -1
)

type Board struct {
	tokens [8][8]int
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
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.tokens[i][j] == WHITE {
				white += 1
			}
		}
	}
	return white, 0
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
