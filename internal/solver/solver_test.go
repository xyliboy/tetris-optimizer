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

func mustParse(t *testing.T, input string) []tetromino.Tetromino {
	t.Helper()
	pieces, err := tetromino.Parse(input)
	if err != nil {
		t.Fatalf("test setup could not parse pieces: %v", err)
	}
	return pieces
}
