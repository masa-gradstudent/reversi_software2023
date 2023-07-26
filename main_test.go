package main

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	b := create_board()
	expected := [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	if !reflect.DeepEqual(b.tokens, expected) {
		t.Errorf("Create Error")
	}

}
func TestDefalut(t *testing.T) {
	b := defalut_board()
	expected := [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, -1, 0, 0, 0},
		{0, 0, 0, -1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	if !reflect.DeepEqual(b.tokens, expected) {
		t.Errorf("Create Error")

	}
}

func TestCount01(t *testing.T) {
	b := create_board()
	white, black := b.count()
	if white != 0 || black != 0 {
		t.Errorf("Count01 Error")
	}
}

func TestCount02(t *testing.T) {
	b := create_board()
	b.tokens[0][0] = WHITE
	b.tokens[0][1] = WHITE
	b.tokens[0][2] = WHITE
	white, black := b.count()
	if white != 3 || black != 0 {
		t.Errorf("Count02 Error")
	}
}

func TestCount03(t *testing.T) {
	b := create_board()
	b.tokens[3][0] = BLACK
	b.tokens[3][1] = BLACK
	b.tokens[3][2] = BLACK
	white, black := b.count()
	if white != 0 || black != 3 {
		t.Errorf("Count03 Error")
	}
}

func TestFlip(t *testing.T) {
	b := create_board()
	b.tokens[2][3] = BLACK
	b.tokens[3][3] = WHITE
	b.tokens[4][3] = WHITE
	b.tokens[5][3] = WHITE
	b.tokens[6][3] = BLACK
	b.flip(6, 3)
	expected := [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, BLACK, 0, 0, 0, 0},
		{0, 0, 0, BLACK, 0, 0, 0, 0},
		{0, 0, 0, BLACK, 0, 0, 0, 0},
		{0, 0, 0, BLACK, 0, 0, 0, 0},
		{0, 0, 0, BLACK, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	if !reflect.DeepEqual(b.tokens, expected) {
		t.Errorf("Flip Error: %d", b.tokens)
	}
}
