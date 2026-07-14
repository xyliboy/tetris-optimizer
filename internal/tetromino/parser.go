package tetromino

import (
	"errors"
	"strings"
)

var ErrInvalidFormat = errors.New("invalid tetromino format")

// Parse converts the complete input file into labeled tetrominoes.
func Parse(data string) ([]Tetromino, error) {
	data = strings.ReplaceAll(data, "\r\n", "\n")
	data = strings.TrimSuffix(data, "\n")
	if data == "" {
		return nil, ErrInvalidFormat
	}

	grids := strings.Split(data, "\n\n")
	if len(grids) > 26 {
		return nil, ErrInvalidFormat
	}

	pieces := make([]Tetromino, 0, len(grids))
	for index, grid := range grids {
		blocks, err := parseGrid(grid)
		if err != nil {
			return nil, err
		}
		pieces = append(pieces, New(rune('A'+index), blocks))
	}
	return pieces, nil
}

func parseGrid(grid string) ([BlockCount]Point, error) {
	var blocks [BlockCount]Point
	lines := strings.Split(grid, "\n")
	if len(lines) != GridSize {
		return blocks, ErrInvalidFormat
	}

	count := 0
	for row, line := range lines {
		if len(line) != GridSize {
			return blocks, ErrInvalidFormat
		}
		for col, cell := range line {
			switch cell {
			case '#':
				if count == BlockCount {
					return blocks, ErrInvalidFormat
				}
				blocks[count] = Point{Row: row, Col: col}
				count++
			case '.':
			default:
				return blocks, ErrInvalidFormat
			}
		}
	}
	if count != BlockCount {
		return blocks, ErrInvalidFormat
	}
	if !isConnected(blocks) {
		return blocks, ErrInvalidFormat
	}
	return blocks, nil
}

func isConnected(blocks [BlockCount]Point) bool {
	occupied := make(map[Point]bool, BlockCount)
	for _, block := range blocks {
		occupied[block] = true
	}

	seen := map[Point]bool{blocks[0]: true}
	queue := []Point{blocks[0]}
	directions := [...]Point{{Row: -1}, {Row: 1}, {Col: -1}, {Col: 1}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			next := Point{Row: current.Row + direction.Row, Col: current.Col + direction.Col}
			if occupied[next] && !seen[next] {
				seen[next] = true
				queue = append(queue, next)
			}
		}
	}
	return len(seen) == BlockCount
}
