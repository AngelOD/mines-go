package core

// Click handles "click"-style events for cells, revealing hidden fields,
// including cascading reveals.
func (board *Board) Click(col, row int) (numChanges int) {
	numChanges = 0

	if board.cells[board.GetIndexFromLocation(col, row)].isRevealed {
		return
	}

	numChanges = board.revealCell(col, row)

	return
}

func (board *Board) GetBoard() []Cell {
	for _, curCell := range board.cells {
		board.findCellType(curCell.Col, curCell.Row)
	}

	return board.cells
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
