package game

import "fmt"

//PLAYER represents player's symbol
const PLAYER = "X"

//OPPONENT represents opponent's (can be AI or other player) symbol
const OPPONENT = "O"

//NEUTRAL represents empty spot on the board
const NEUTRAL = "_"

//Board represents the game board
//|1|2|3|
//|4|5|6|
//|7|8|9|
type Board struct {
	Field [3][3]string
}

//Make (method of Board) -> Initializing the start board with 9 empty spots
//|_|_|_|
//|_|_|_|
//|_|_|_|
func (b *Board) Make() {
	b.Field = [3][3]string{{NEUTRAL, NEUTRAL, NEUTRAL},
		{NEUTRAL, NEUTRAL, NEUTRAL},
		{NEUTRAL, NEUTRAL, NEUTRAL}}
}

//SetPossition -> Changing a position on coordinates i (for row) and j (or column)
//with the given symbol (symbol)
func (b *Board) SetPossition(symbol string, i, j int) {
	b.Field[i][j] = symbol
}

//PrintBoard -> Printing the board
func (b Board) PrintBoard() {
	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Print("|")
		for j := 0; j < 3; j++ {
			fmt.Printf("%s%s", b.Field[i][j], "|")
		}
		fmt.Println()
	}
}

//IsFull -> Check for full board
func (b Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.PositionIsEmpty(i, j) {
				return true
			}
		}
	}
	return false
}

//Evaluate -> Checking for winner
func (b Board) Evaluate() int {
	for row := 0; row < 3; row++ {
		if b.Field[row][0] == b.Field[row][1] &&
			b.Field[row][1] == b.Field[row][2] {
			if b.Field[row][0] == PLAYER {
				return 10
			} else if b.Field[row][0] == OPPONENT {
				return -10
			}
		}
	}

	for col := 0; col < 3; col++ {
		if b.Field[0][col] == b.Field[1][col] &&
			b.Field[1][col] == b.Field[2][col] {
			if b.Field[0][col] == PLAYER {
				return 10
			} else if b.Field[0][col] == OPPONENT {
				return -10
			}
		}
	}

	if b.Field[0][0] == b.Field[1][1] &&
		b.Field[1][1] == b.Field[2][2] {
		if b.Field[0][0] == PLAYER {
			return 10
		} else if b.Field[0][0] == OPPONENT {
			return -10
		}
	}

	if b.Field[0][2] == b.Field[1][1] &&
		b.Field[1][1] == b.Field[2][0] {
		if b.Field[0][2] == PLAYER {
			return 10
		} else if b.Field[0][2] == OPPONENT {
			return -10
		}
	}
	return 0
}

//PositionIsEmpty -> Check for <i> <j> position is empty or not
func (b Board) PositionIsEmpty(i, j int) bool {
	return b.Field[i][j] == NEUTRAL
}

//CheckForIdenticalSymbolsInRow -> Check for winner on row
func (b Board) CheckForIdenticalSymbolsInRow(row int) bool {
	return b.Field[row][0] == b.Field[row][1] && b.Field[row][1] == b.Field[row][2]
}

//CheckForIdenticalSymbolsInCol -> Check for winner on column
func (b Board) CheckForIdenticalSymbolsInCol(col int) bool {
	return b.Field[0][col] == b.Field[1][col] && b.Field[1][col] == b.Field[2][col]
}

//CheckForWinnerInFirstDiagonal -> Check for winner on main diagonal
func (b Board) CheckForWinnerInFirstDiagonal() bool {
	return b.Field[0][0] == b.Field[1][1] && b.Field[1][1] == b.Field[2][2]
}

//CheckForWinnerInSecondDiagonal -> Check for winner on secondary diagonal
func (b Board) CheckForWinnerInSecondDiagonal() bool {
	return b.Field[0][2] == b.Field[1][1] && b.Field[1][1] == b.Field[2][0]
}

//GetWinnerValue -> Used for alpha-beta algorithm (heuristic function)
func GetWinnerValue(symbol string) int {
	if symbol == PLAYER {
		return 10
	} else if symbol == OPPONENT {
		return -10
	}
	return 0
}
