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
func TestDetermineWinner(t *testing.T) {
	white1, black1 := 20, 15
	result1 := determineWinner(white1, black1)
	if result1 != "White wins!" {
		t.Errorf("Test case 1 failed: Expected 'White wins!', got '%s'", result1)
	}

	white2, black2 := 12, 18
	result2 := determineWinner(white2, black2)
	if result2 != "Black wins!" {
		t.Errorf("Test case 2 failed: Expected 'Black wins!', got '%s'", result2)
	}

	white3, black3 := 30, 30
	result3 := determineWinner(white3, black3)
	if result3 != "It's a tie!" {
		t.Errorf("Test case 3 failed: Expected 'It's a tie!', got '%s'", result3)
	}
}
func TestSwitchTurn(t *testing.T) {
	b := create_board()
	b.currentTurn = WHITE
	b.switchTurn()
	if b.currentTurn != BLACK {
		t.Errorf("TestSwitchTurn failed: Expected BLACK as the current player, got %d", b.currentTurn)
	}

	b.switchTurn()
	if b.currentTurn != WHITE {
		t.Errorf("TestSwitchTurn failed: Expected WHITE as the current player, got %d", b.currentTurn)
	}
}
