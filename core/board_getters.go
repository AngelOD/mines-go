package core

// ===== Public =====

func (board *Board) GetGameStatus() GameStatus {
	if board.gameStatus == GAME_RUNNING {
		var numHiddenFree int

		for _, curCell := range board.cells {
			if !curCell.isRevealed && !curCell.hasMine {
				numHiddenFree++
			}
		}

		if numHiddenFree == 0 {
			board.gameStatus = GAME_WON
		}
	}

	return board.gameStatus
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
