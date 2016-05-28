package core

// Click takes care of revealing fields that are hidden,
func (board *Board) Click(col, row int) (err error) {
	if board.cells[board.GetIndexFromLocation(col, row)].isRevealed {
		err = nil
		return
	}

	return
}

func (board *Board) GetBoard() (cells []Cell) {
	var cellNum int
	var cellType CellType

	cells = make([]Cell, 0)

	for _, curCell := range board.cells {
		cellType, cellNum = board.getCellType(curCell.locCol, curCell.locRow)
		cells = append(cells, newCell(curCell.locCol, curCell.locRow, cellNum, cellType))
	}

	return
}

func (board *Board) GetIndexFromLocation(col, row int) (index int) {
	index = row*board.colCount + col

	return
}

func (board *Board) GetLocationFromIndex(index int) (col, row int) {
	col = index % board.colCount
	row = index / board.colCount

	return
}
