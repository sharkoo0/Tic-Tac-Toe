package game

import "fmt"

//Move represents the coordinates of the last move
type Move struct {
	Row, Col int
}

//Make (method of Move) -> setting default value of Move
func (m *Move) Make() {
	m.Row = -1
	m.Col = -1
}

//PrintMove -> printing the Move
func (m Move) PrintMove() {
	fmt.Printf("%d , %d", m.Row, m.Col)
}
