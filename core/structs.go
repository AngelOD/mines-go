package core

type cell struct {
	board              *Board
	locCol, locRow     int
	hasMine            bool
	isRevealed         bool
	proximityMineCount int
}

// Board manages the Mines-Go game, from board to rules.
type Board struct {
	cells              []cell
	rowCount, colCount int
	seed               int64
}

// NewBoard instantiates a new Board and returns a pointer to it.
func NewBoard(numCols, numRows, numMines int) (board *Board) {
	board = new(Board)
	board.init(numCols, numRows, numMines)

	return
}

func newCell(board *Board, locCol, locRow int) (retCell cell) {
	retCell.board = board
	retCell.locCol = locCol
	retCell.locRow = locRow

	return
}
