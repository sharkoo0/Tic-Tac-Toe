package algorithms

import "github.com/sharkoo0/Tic-Tac-Toe/pkg/game"
import "math"

//AlphaBeta -> The alpha-beta algorithm 
//Parameters: <board>, <depth>, <alpha>, <beta>, <isMax>
//<board> -> The gaming board for this game
//<depth> -> The current depth of the solutions tree
//<alpha>, <beta> -> represent the maximizer and the minimizer respectivly
//<isMax> -> Checking whos turn is (maximizer(alpha) or minimizer(beta))
func AlphaBeta(board game.Board, depth int, alpha int, beta int, isMaximizer bool) int {
	result := board.Evaluate()

	if result == 10 {
		return result - depth
	}

	if result == -10 {
		return result + depth
	}

	if !board.IsFull() {
		return 0
	}

	if isMaximizer {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board.PositionIsEmpty(i, j) {
					board.SetPossition(game.PLAYER, i, j)
					alpha = int(math.Max(float64(alpha), float64(AlphaBeta(board, depth+1, alpha, beta, !isMaximizer))))
					board.SetPossition(game.NEUTRAL, i, j)
					if alpha >= beta {
						break
					}
				}
			}
		}
		return alpha
	} 
	//else statement
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.PositionIsEmpty(i, j) {
				board.SetPossition(game.OPPONENT, i, j)				
				beta = int(math.Min(float64(beta), float64(AlphaBeta(board, depth+1, alpha, beta, !isMaximizer))))
				board.SetPossition(game.NEUTRAL, i, j)
				if alpha >= beta {
					break
				}
			}
		}
	}
	return beta
	
}

//BestMove -> This function is used for AI player. This is the AI's turn.
func BestMove(board game.Board, isMaximizer bool) game.Move {
	alpha := -1000
	beta := 1000
	bestValue := 0
	if isMaximizer {
		bestValue = alpha
	} else {
		bestValue = beta
	}
	var bestMove game.Move
	bestMove.Make()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.PositionIsEmpty(i, j) {
				if isMaximizer {
					board.SetPossition(game.PLAYER, i, j)
				} else {
					board.SetPossition(game.OPPONENT, i, j)
				}
				value := AlphaBeta(board, 0, alpha, beta, !isMaximizer)
				board.SetPossition(game.NEUTRAL, i, j)
				if isMaximizer {
					if value > bestValue {
						bestMove.Row = i
						bestMove.Col = j
						bestValue = value
					}
				} else {
					if value < bestValue {
						bestMove.Row = i
						bestMove.Col = j
						bestValue = value
					}
				}
			}
		}
	}

	return bestMove
}