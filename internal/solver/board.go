package solver

import (
	"strings"

	"tetris-optimizer/internal/tetromino"
)

// Board is a square grid used while searching for a solution.
type Board struct {
	size  int
	cells [][]rune
}

func newBoard(size int) *Board {
	cells := make([][]rune, size)
	for row := range cells {
		cells[row] = make([]rune, size)
		for col := range cells[row] {
			cells[row][col] = '.'
		}
	}
	return &Board{size: size, cells: cells}
}

func (board *Board) canPlace(piece tetromino.Tetromino, row, col int) bool {
	if row+piece.Height > board.size || col+piece.Width > board.size {
		return false
	}
	for _, block := range piece.Blocks {
		if board.cells[row+block.Row][col+block.Col] != '.' {
			return false
		}
	}
	return true
}

func (board *Board) place(piece tetromino.Tetromino, row, col int, value rune) {
	for _, block := range piece.Blocks {
		board.cells[row+block.Row][col+block.Col] = value
	}
}

func (board *Board) String() string {
	var output strings.Builder
	output.Grow(board.size * (board.size + 1))
	for row, cells := range board.cells {
		output.WriteString(string(cells))
		if row < board.size-1 {
			output.WriteByte('\n')
		}
	}
	return output.String()
}
