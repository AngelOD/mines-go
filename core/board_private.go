package core

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func (board *Board) findProximityMineCount(col, row int) {
	var mineCount int

	cells := board.findSurroundingCells(col, row)

	for _, curCell := range cells {
		if curCell.hasMine {
			mineCount++
		}
	}

	board.cells[board.GetIndexFromLocation(col, row)].proximityMineCount = mineCount
}

func (board *Board) findSurroundingCells(col, row int) []*internalCell {
	var cells []*internalCell

	for curCol := col - 1; curCol <= col+1; curCol++ {
		for curRow := row - 1; curRow <= row+1; curRow++ {
			if (curCol == col && curRow == row) || !board.isValidCoord(curCol, curRow) {
				continue
			}

			cells = append(cells, &board.cells[board.GetIndexFromLocation(curCol, curRow)])
		}
	}

	return cells
}

func (board *Board) getCellType(col, row int) (cellType CellType, cellNum int) {
	curCell := board.cells[board.GetIndexFromLocation(col, row)]

	cellType = NOT_REVEALED
	cellNum = -1

	if curCell.isRevealed {
		if curCell.hasMine {
			cellType = HAS_MINE
		} else {
			cellType = HAS_NUMBER
			cellNum = curCell.proximityMineCount
		}
	}

	return
}

func (board *Board) init(numCols, numRows, numMines int) (err error) {
	if numCols < 10 || numRows < 10 {
		err = fmt.Errorf("Grid must be at least 10x10.")
		return
	}

	maxMines := int(math.Floor((float64(numCols) * float64(numRows)) * float64(0.9)))

	if numMines < 1 || numMines > maxMines {
		err = fmt.Errorf("There must be between 1 and %d mines for a grid of this size.", maxMines)
		return
	}

	// Instantiate board and all cells
	board.colCount = numCols
	board.rowCount = numRows

	board.initCells()

	// Fill with desired number of mines
	board.placeMines(numMines)

	return
}

func (board *Board) initCells() {
	var cells []internalCell
	var curCol, curRow int

	for {
		cells = append(cells, newInternalCell(board, curCol, curRow))

		curCol++
		if curCol == board.colCount {
			curCol = 0

			curRow++
			if curRow == board.rowCount {
				break
			}
		}
	}

	board.cells = cells
}

func (board *Board) isValidCoord(col, row int) bool {
	if col < 0 || row < 0 || col >= board.colCount || row >= board.rowCount {
		return false
	}

	return true
}

func (board *Board) placeMines(mineCount int) {
	var locs []int
	var minesPlaced int

	// Generate pseudo-set of possible locations
	for i := 0; i < board.colCount; i++ {
		for j := 0; j < board.rowCount; j++ {
			locs = append(locs, board.GetIndexFromLocation(i, j))
		}
	}

	// Seed the randomizer
	board.seed = time.Now().UnixNano()
	rand.Seed(board.seed)

	for minesPlaced < mineCount {
		numLocs := len(locs)
		i := rand.Intn(numLocs)
		loc := locs[i]

		board.cells[loc].hasMine = true

		locs = append(locs[0:i], locs[i+1:]...)

		minesPlaced++
	}
}

func (board *Board) revealCell(col, row int) {
	curCell := board.cells[board.GetIndexFromLocation(col, row)]

	if curCell.isRevealed {
		return
	}

	curCell.isRevealed = true

	if curCell.hasMine {
		// TODO Consider how to handle this better
		return
	}

	board.findProximityMineCount(col, row)

	if curCell.proximityMineCount > 0 {
		return
	}

	surroundingCells := board.findSurroundingCells(col, row)

	for _, otherCell := range surroundingCells {
		board.revealCell(otherCell.locCol, otherCell.locRow)
	}
}
