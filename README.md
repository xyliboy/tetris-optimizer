# Tetris Optimizer

Tetris Optimizer reads tetrominoes from a text file and arranges them in the
smallest possible square. Pieces keep their input order and are displayed as
uppercase letters (`A`, `B`, `C`, ...). Empty cells are displayed as `.`.

## Run

```sh
go run . path/to/pieces.txt
```

Invalid arguments, unreadable files, and malformed pieces produce:

```text
ERROR
```

## Input format

Each piece is a 4-by-4 grid containing exactly four connected `#` cells.
All other cells must be `.`. Separate consecutive pieces with one empty line.

```text
#...
#...
#...
#...

....
.##.
.##.
....
```

Blocks connect horizontally or vertically; diagonal contact alone is invalid.
The program accepts both Unix (`LF`) and Windows (`CRLF`) line endings and a
maximum of 26 pieces.

## Project map

```text
main.go                         command-line input and output
internal/tetromino/parser.go    file-format parsing and validation
internal/tetromino/tetromino.go piece model and normalization
internal/solver/board.go        board placement operations
internal/solver/solver.go       smallest-square backtracking search
```

This separation follows the program's data flow: parse and validate the input,
normalize each piece, search for the smallest arrangement, then print it.

## Development

Run all tests:

```sh
go test ./...
```

Format and statically check the project:

```sh
gofmt -w .
go vet ./...
```

The project uses only the Go standard library.
