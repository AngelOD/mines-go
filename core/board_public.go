package core

// Click handles "click"-style events for cells, revealing hidden fields,
// including cascading reveals.
func (board *Board) Click(col, row int) (numChanges int, gameStatus GameStatus) {
	numChanges = 0
	gameStatus = board.gameStatus

	if board.cells[board.GetIndexFromLocation(col, row)].isRevealed {
		return
	}

	numChanges = board.revealCell(col, row)
	gameStatus = board.GetGameStatus()

	return
}
