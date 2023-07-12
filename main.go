package main

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
	b.tokens[3][3] = 1
	b.tokens[3][4] = -1
	b.tokens[4][3] = -1
	b.tokens[4][4] = 1
	return b
}
