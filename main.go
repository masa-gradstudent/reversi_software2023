package main

type Board struct {
	tokens [8][8]int
}

func create_board() Board {
	b := Board{}
	b.tokens = [8][8]int{}
	return b
}
