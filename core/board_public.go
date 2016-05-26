package core

// Click takes care of revealing fields that are hidden,
func (board *Board) Click(col, row int) (err error) {
	if board.cells[col][row].isRevealed {
		err = nil
		return
	}

	return
}
