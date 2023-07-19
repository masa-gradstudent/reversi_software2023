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
