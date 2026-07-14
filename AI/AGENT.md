# Agent Guide

## Project goal

Build a Go command-line program that reads tetrominoes from one file and places
them in the smallest possible square. The first piece is printed as `A`, the
second as `B`, and so on. Empty cells are printed as `.`.

The program is run with:

```sh
go run . path/to/pieces.txt
```

For every invalid argument, unreadable file, or malformed input, print exactly:

```text
ERROR
```

Do not add extra user-facing output.

## Repository map

- `main.go`: validates command-line usage, reads the file, and prints the result.
- `internal/tetromino/`: owns the piece model, normalization, parsing, and input validation.
- `internal/solver/`: owns the square board and the backtracking search.
- `AI/`: contains the original assignment and audit material. Treat these files as reference, not application code.
- `README.md`: explains installation, usage, input format, architecture, and development commands.

## Input rules

- Accept between 1 and 26 tetrominoes.
- Each tetromino is exactly four rows of four characters.
- Only `.` and `#` are valid characters.
- Each tetromino contains exactly four `#` cells.
- Blocks must connect through horizontal or vertical edges.
- Diagonal contact does not count as a connection.
- Consecutive tetrominoes are separated by one empty line.
- Accept Unix (`LF`) and Windows (`CRLF`) line endings.
- A final newline is optional.
- Do not rotate or mirror pieces.

## Design rules

- Keep parsing and validation in `internal/tetromino`.
- Keep placement and search logic in `internal/solver`.
- Keep `main.go` small; it should only coordinate input, solving, and output.
- Prefer descriptive names and short functions with one responsibility.
- Use only the Go standard library.
- Avoid allocations inside the recursive search unless they materially simplify correctness.
- Preserve piece order and its corresponding output label.
- Return errors from internal code; only the CLI decides to print `ERROR`.

## Required checks

Before considering a change complete, run:

```sh
gofmt -w .
go test ./...
go vet ./...
git diff --check
```

Add or update tests whenever behavior changes. Tests should cover both the
successful path and malformed-input edge cases.

## Definition of done

The project is ready only when:

1. It compiles successfully.
2. All tests and static checks pass.
3. Known good and bad evaluator examples behave correctly.
4. Every solution uses the smallest possible square.
5. The hard example completes within the evaluator's time limit.
6. The repository contains no debug output or temporary files.
7. The README matches the final behavior.

Do not remove `AI/info.txt` or `AI/audit.txt` during normal development. Remove
them only after the user explicitly confirms that the project is ready for its
final cleanup.
