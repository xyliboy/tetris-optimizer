package unit_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLIPrintsSolvedBoard(t *testing.T) {
	input := writeInput(t, squarePiece+"\n")
	output := runCLI(t, input)
	if got := strings.TrimSpace(output); got != "AA\nAA" {
		t.Errorf("CLI output = %q, want %q", got, "AA\nAA")
	}
}

func TestCLIPrintsErrorForInvalidUsage(t *testing.T) {
	tests := [][]string{
		nil,
		{"one", "two"},
		{filepath.Join(t.TempDir(), "missing.txt")},
		{writeInput(t, "not a tetromino")},
		{writeInput(t, "")},
		{t.TempDir()},
	}
	for _, args := range tests {
		if got := strings.TrimSpace(runCLI(t, args...)); got != "ERROR" {
			t.Errorf("CLI output for %v = %q, want ERROR", args, got)
		}
	}
}

func TestAllTetrominoFixture(t *testing.T) {
	output := runCLI(t, filepath.Join("..", "examples", "all_tetrominoes.txt"))
	for label := 'A'; label <= 'G'; label++ {
		if strings.Count(output, string(label)) != tetrominoBlockCount {
			t.Errorf("fixture output contains the wrong number of %c cells:\n%s", label, output)
		}
	}
}

const tetrominoBlockCount = 4

func runCLI(t *testing.T, args ...string) string {
	t.Helper()
	commandArgs := append([]string{"run", ".."}, args...)
	command := exec.Command("go", commandArgs...)
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\n%s", err, output)
	}
	return string(output)
}

func writeInput(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "pieces.txt")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("could not create test input: %v", err)
	}
	return path
}
