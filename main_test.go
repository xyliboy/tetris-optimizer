package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunRejectsWrongArgumentCount(t *testing.T) {
	for _, args := range [][]string{nil, {"one", "two"}} {
		if err := run(args, &bytes.Buffer{}); err == nil {
			t.Errorf("run(%v) returned no error", args)
		}
	}
}

func TestRunRejectsUnreadableOrMalformedFiles(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing.txt")
	if err := run([]string{missing}, &bytes.Buffer{}); err == nil {
		t.Error("run() accepted a missing file")
	}

	invalid := writeInput(t, "not a tetromino")
	if err := run([]string{invalid}, &bytes.Buffer{}); err == nil {
		t.Error("run() accepted malformed input")
	}

	empty := writeInput(t, "")
	if err := run([]string{empty}, &bytes.Buffer{}); err == nil {
		t.Error("run() accepted an empty file")
	}

	if err := run([]string{t.TempDir()}, &bytes.Buffer{}); err == nil {
		t.Error("run() accepted a directory as input")
	}
}

func TestRunPrintsSolvedBoard(t *testing.T) {
	input := writeInput(t, "##..\n##..\n....\n....\n")
	var output bytes.Buffer
	if err := run([]string{input}, &output); err != nil {
		t.Fatalf("run() returned an unexpected error: %v", err)
	}
	if got := strings.TrimSpace(output.String()); got != "AA\nAA" {
		t.Errorf("run() output = %q, want %q", got, "AA\nAA")
	}
}

func writeInput(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "pieces.txt")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("could not create test input: %v", err)
	}
	return path
}
