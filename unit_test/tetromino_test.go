package unit_test

import (
	"strings"
	"testing"

	"tetris-optimizer/internal/tetromino"
)

const squarePiece = "##..\n##..\n....\n...."

func TestParseValidPieces(t *testing.T) {
	input := "#...\n#...\n#...\n#...\n\n....\n.##.\n.##.\n....\n"
	pieces, err := tetromino.Parse(input)
	if err != nil {
		t.Fatalf("Parse() returned an unexpected error: %v", err)
	}
	if len(pieces) != 2 || pieces[0].Label != 'A' || pieces[1].Label != 'B' {
		t.Fatalf("Parse() returned incorrectly labeled pieces: %+v", pieces)
	}
	if pieces[0].Width != 1 || pieces[0].Height != 4 || pieces[1].Width != 2 || pieces[1].Height != 2 {
		t.Errorf("Parse() did not normalize piece dimensions: %+v", pieces)
	}
}

func TestParseAcceptsCommonLineEndings(t *testing.T) {
	inputs := []string{
		"##..\r\n##..\r\n....\r\n....\r\n",
		"##..\r\n##..\n....\r\n....\n",
	}
	for _, input := range inputs {
		if _, err := tetromino.Parse(input); err != nil {
			t.Errorf("Parse() rejected valid line endings: %v", err)
		}
	}
}

func TestParseAcceptsBoundaryPieceCount(t *testing.T) {
	pieces, err := tetromino.Parse(strings.Repeat(squarePiece+"\n\n", 25) + squarePiece)
	if err != nil {
		t.Fatalf("Parse() rejected 26 pieces: %v", err)
	}
	if len(pieces) != 26 || pieces[25].Label != 'Z' {
		t.Fatalf("Parse() did not label the boundary piece as Z: %+v", pieces[25])
	}
	if _, err := tetromino.Parse(strings.Repeat(squarePiece+"\n\n", 26) + squarePiece); err == nil {
		t.Fatal("Parse() accepted more labels than A-Z can represent")
	}
}

func TestParseRejectsInvalidInput(t *testing.T) {
	tests := map[string]string{
		"empty file":          "",
		"too few rows":        "####\n....\n....",
		"wrong row width":     "####.\n....\n....\n....",
		"invalid character":   "###x\n....\n....\n....",
		"too few blocks":      "###.\n....\n....\n....",
		"too many blocks":     "####\n#...\n....\n....",
		"disconnected blocks": "#.#.\n....\n#.#.\n....",
		"extra separator":     squarePiece + "\n\n\n" + squarePiece,
		"leading newline":     "\n" + squarePiece,
		"trailing blank line": squarePiece + "\n\n",
		"space separator":     squarePiece + "\n \n" + squarePiece,
		"spaces in grid":      "##  \n##..\n....\n....",
		"tab in grid":         "##.\t\n##..\n....\n....",
		"carriage returns":    "##..\r##..\r....\r....\r",
		"diagonal chain":      "#...\n.#..\n..#.\n...#",
		"two separate pairs":  "##..\n....\n##..\n....",
		"nul byte":            "##..\n##..\n....\n...\x00",
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := tetromino.Parse(input); err == nil {
				t.Fatal("Parse() accepted invalid input")
			}
		})
	}
}
