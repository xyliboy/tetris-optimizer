package unit_test

import (
	"strings"
	"testing"

	"tetris-optimizer/internal/solver"
	"tetris-optimizer/internal/tetromino"
)

func TestSolveUsesSmallestSquare(t *testing.T) {
	pieces := mustParse(t, strings.Repeat(squarePiece+"\n\n", 3)+squarePiece)
	result := solver.Solve(pieces).String()
	lines := strings.Split(result, "\n")
	if len(lines) != 4 || strings.Contains(result, ".") {
		t.Fatalf("Solve() did not fill a minimal 4x4 board:\n%s", result)
	}
}

func TestSolveRespectsPieceDimensions(t *testing.T) {
	tests := map[string]struct {
		input string
		size  int
	}{
		"square":          {squarePiece, 2},
		"horizontal line": {"####\n....\n....\n....", 4},
		"vertical line":   {"#...\n#...\n#...\n#...", 4},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lines := strings.Split(solver.Solve(mustParse(t, test.input)).String(), "\n")
			if len(lines) != test.size || len(lines[0]) != test.size {
				t.Fatalf("board is %dx%d, want %dx%d", len(lines), len(lines[0]), test.size, test.size)
			}
		})
	}
}

func TestSolvePreservesEveryPiece(t *testing.T) {
	pieces := mustParse(t, ".##.\n##..\n....\n....\n\n.#..\n###.\n....\n....\n\n####\n....\n....\n....")
	result := solver.Solve(pieces).String()
	lines := strings.Split(result, "\n")
	for row, line := range lines {
		if len(line) != len(lines) {
			t.Fatalf("row %d has width %d on a board with height %d", row, len(line), len(lines))
		}
	}
	for _, piece := range pieces {
		if strings.Count(result, string(piece.Label)) != tetromino.BlockCount {
			t.Fatalf("output does not contain four %c cells:\n%s", piece.Label, result)
		}
		var found [tetromino.BlockCount]tetromino.Point
		index := 0
		for row, line := range lines {
			for col, cell := range line {
				if cell == piece.Label {
					found[index] = tetromino.Point{Row: row, Col: col}
					index++
				}
			}
		}
		if placed := tetromino.New(piece.Label, found); placed.Blocks != piece.Blocks {
			t.Errorf("piece %c changed shape: got %v, want %v", piece.Label, placed.Blocks, piece.Blocks)
		}
	}
}

func mustParse(t *testing.T, input string) []tetromino.Tetromino {
	t.Helper()
	pieces, err := tetromino.Parse(input)
	if err != nil {
		t.Fatalf("test setup could not parse pieces: %v", err)
	}
	return pieces
}
