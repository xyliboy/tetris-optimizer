package tetromino

import "testing"

func TestParseValidPieces(t *testing.T) {
	input := "#...\n#...\n#...\n#...\n\n....\n.##.\n.##.\n....\n"
	pieces, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse() returned an unexpected error: %v", err)
	}
	if len(pieces) != 2 {
		t.Fatalf("Parse() returned %d pieces, want 2", len(pieces))
	}
	if pieces[0].Label != 'A' || pieces[0].Width != 1 || pieces[0].Height != 4 {
		t.Errorf("first piece was not labeled or normalized correctly: %+v", pieces[0])
	}
	if pieces[1].Label != 'B' || pieces[1].Width != 2 || pieces[1].Height != 2 {
		t.Errorf("second piece was not labeled or normalized correctly: %+v", pieces[1])
	}
}

func TestParseAcceptsWindowsLineEndings(t *testing.T) {
	_, err := Parse("##..\r\n##..\r\n....\r\n....\r\n")
	if err != nil {
		t.Fatalf("Parse() rejected CRLF input: %v", err)
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
		"extra separator":     "##..\n##..\n....\n....\n\n\n##..\n##..\n....\n....",
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := Parse(input); err == nil {
				t.Fatal("Parse() accepted invalid input")
			}
		})
	}
}
