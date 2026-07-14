package solver

import (
	"strings"
	"testing"

	"tetris-optimizer/internal/tetromino"
)

func TestSolveUsesSmallestSquare(t *testing.T) {
	pieces := mustParse(t, "##..\n##..\n....\n....\n\n##..\n##..\n....\n....\n\n##..\n##..\n....\n....\n\n##..\n##..\n....\n....")
	board := Solve(pieces)

	lines := strings.Split(board.String(), "\n")
	if len(lines) != 4 {
		t.Fatalf("Solve() produced a %dx%d board, want 4x4", len(lines), len(lines))
	}
	if strings.Contains(board.String(), ".") {
		t.Errorf("Solve() left empty cells:\n%s", board)
	}
}

func TestSolveKeepsEveryPieceLabel(t *testing.T) {
	pieces := mustParse(t, "#...\n#...\n#...\n#...\n\n####\n....\n....\n....")
	result := Solve(pieces).String()
	for _, label := range []rune{'A', 'B'} {
		if strings.Count(result, string(label)) != tetromino.BlockCount {
			t.Errorf("output contains the wrong number of %c blocks:\n%s", label, result)
		}
	}
}

func TestSolveRespectsPieceDimensions(t *testing.T) {
	tests := map[string]struct {
		input string
		size  int
	}{
		"square fits two by two":             {"##..\n##..\n....\n....", 2},
		"horizontal line needs four columns": {"####\n....\n....\n....", 4},
		"vertical line needs four rows":      {"#...\n#...\n#...\n#...", 4},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			board := Solve(mustParse(t, test.input))
			lines := strings.Split(board.String(), "\n")
			if len(lines) != test.size || len(lines[0]) != test.size {
				t.Fatalf("board dimensions are %dx%d, want %dx%d", len(lines), len(lines[0]), test.size, test.size)
			}
		})
	}
}

func TestSolveOutputInvariants(t *testing.T) {
	pieces := mustParse(t, ".##.\n##..\n....\n....\n\n.#..\n###.\n....\n....\n\n####\n....\n....\n....")
	result := Solve(pieces).String()
	lines := strings.Split(result, "\n")
	for row, line := range lines {
		if len(line) != len(lines) {
			t.Fatalf("row %d has width %d on a board with height %d", row, len(line), len(lines))
		}
		for _, cell := range line {
			if cell != '.' && (cell < 'A' || cell > 'C') {
				t.Fatalf("output contains unexpected character %q", cell)
			}
		}
	}
	for _, piece := range pieces {
		if strings.Count(result, string(piece.Label)) != tetromino.BlockCount {
			t.Errorf("output does not contain exactly four %c cells:\n%s", piece.Label, result)
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
		placed := tetromino.New(piece.Label, found)
		if placed.Blocks != piece.Blocks {
			t.Errorf("piece %c was rotated, mirrored, or distorted: got %v, want %v", piece.Label, placed.Blocks, piece.Blocks)
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
