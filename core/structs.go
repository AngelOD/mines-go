package core

type Cell struct {
	Col   int
	Row   int
	CType CellType
	CNum  int

	board              *Board
	checked            bool
	hasMine            bool
	isRevealed         bool
	proximityMineCount int
}

// Board manages the Mines-Go game, from board to rules.
type Board struct {
	cells              []Cell
	rowCount, colCount int
	seed               int64
	gameStatus         GameStatus
}

// NewBoard instantiates a new Board with a random seed and returns a pointer
// to it.
func NewBoard(numCols, numRows, numMines int) (board *Board) {
	board = new(Board)
	board.init(numCols, numRows, numMines)

	return
}

// NewBoardWithSeed instantiates a new Board with a specific seed and returns a
// pointer to it.
func NewBoardWithSeed(numCols, numRows, numMines int, seed int64) (board *Board) {
	board = new(Board)
	board.seed = seed
	board.init(numCols, numRows, numMines)

	return
}

func newInternalCell(board *Board, col, row int) (retCell Cell) {
	retCell.board = board
	retCell.Col = col
	retCell.Row = row

	return
}

func newCell(col, row, cellNum int, cellType CellType) (retCell Cell) {
	retCell.Col = col
	retCell.Row = row
	retCell.CType = cellType
	retCell.CNum = cellNum

	return
}
