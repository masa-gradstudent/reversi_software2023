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
		t.Errorf("Count Error")
	}
}

func TestCount02(t *testing.T) {
	b := create_board()
	b.tokens[0][0] = WHITE
	b.tokens[0][1] = WHITE
	b.tokens[0][2] = WHITE
	white, black := b.count()
	if white != 3 || black != 0 {
		t.Errorf("Count Error")
	}
}
