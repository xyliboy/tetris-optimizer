package solver

import "tetris-optimizer/internal/tetromino"

// Solve returns the smallest square that can contain every piece in order.
func Solve(pieces []tetromino.Tetromino) *Board {
	size := smallestPossibleSize(len(pieces) * tetromino.BlockCount)
	for {
		board := newBoard(size)
		if arrange(board, pieces, 0) {
			return board
		}
		size++
	}
}

func smallestPossibleSize(area int) int {
	size := 1
	for size*size < area {
		size++
	}
	return size
}

func arrange(board *Board, pieces []tetromino.Tetromino, index int) bool {
	if index == len(pieces) {
		return true
	}

	piece := pieces[index]
	for row := 0; row <= board.size-piece.Height; row++ {
		for col := 0; col <= board.size-piece.Width; col++ {
			if !board.canPlace(piece, row, col) {
				continue
			}
			board.place(piece, row, col, piece.Label)
			if arrange(board, pieces, index+1) {
				return true
			}
			board.place(piece, row, col, '.')
		}
	}
	return false
}
