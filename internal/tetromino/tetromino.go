package tetromino

const (
	GridSize   = 4
	BlockCount = 4
)

// Point is one occupied cell relative to a tetromino's top-left corner.
type Point struct {
	Row int
	Col int
}

// Tetromino is a normalized, immutable piece identified by its output letter.
type Tetromino struct {
	Label  rune
	Blocks [BlockCount]Point
	Width  int
	Height int
}

// New builds a normalized tetromino from four validated points.
func New(label rune, blocks [BlockCount]Point) Tetromino {
	minRow, minCol := blocks[0].Row, blocks[0].Col
	for _, block := range blocks[1:] {
		if block.Row < minRow {
			minRow = block.Row
		}
		if block.Col < minCol {
			minCol = block.Col
		}
	}

	width, height := 0, 0
	for i := range blocks {
		blocks[i].Row -= minRow
		blocks[i].Col -= minCol
		if blocks[i].Col+1 > width {
			width = blocks[i].Col + 1
		}
		if blocks[i].Row+1 > height {
			height = blocks[i].Row + 1
		}
	}

	return Tetromino{Label: label, Blocks: blocks, Width: width, Height: height}
}
