package core

type Cell struct {
	Col   int
	Row   int
	CType CellType
	CNum  int
}

type internalCell struct {
	board              *Board
	locCol, locRow     int
	hasMine            bool
	isRevealed         bool
	proximityMineCount int
}

// Board manages the Mines-Go game, from board to rules.
type Board struct {
	cells              []internalCell
	rowCount, colCount int
	seed               int64
}

// NewBoard instantiates a new Board and returns a pointer to it.
func NewBoard(numCols, numRows, numMines int) (board *Board) {
	board = new(Board)
	board.init(numCols, numRows, numMines)

	return
}

func newInternalCell(board *Board, locCol, locRow int) (retCell internalCell) {
	retCell.board = board
	retCell.locCol = locCol
	retCell.locRow = locRow

	return
}

func newCell(col, row, cellNum int, cellType CellType) (retCell Cell) {
	retCell.Col = col
	retCell.Row = row
	retCell.CType = cellType
	retCell.CNum = cellNum

	return
}
