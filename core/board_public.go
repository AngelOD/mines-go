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

func (board *Board) MarkCell(col, row int) bool {
	if board.gameStatus != GAME_RUNNING {
		return false
	}

	cellIndex := board.GetIndexFromLocation(col, row)

	if board.cells[cellIndex].isRevealed {
		return false
	}

	board.cells[cellIndex].isMarked = !board.cells[cellIndex].isMarked

	return true
}

func (board *Board) RevealAllMines() bool {
	if board.gameStatus == GAME_RUNNING {
		return false
	}

	for _, curCell := range board.cells {
		if curCell.hasMine {
			curCell.isRevealed = true
			board.findCellType(curCell.Col, curCell.Row)
		}
	}

	return true
}
