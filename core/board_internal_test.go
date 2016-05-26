package core

import (
	"testing"
)

var boardInitTests = []struct {
	colCount  int
	rowCount  int
	mineCount int
}{
	{10, 10, 1},
	{10, 10, 10},
	{10, 10, 50},
	{10, 10, 90},
	{100, 100, 1},
	{100, 100, 9000},
}

func TestBoardInit(t *testing.T) {
	for _, testData := range boardInitTests {
		var numMines int

		board := NewBoard(testData.colCount, testData.rowCount, testData.mineCount)

		for i := 0; i < testData.colCount; i++ {
			for j := 0; j < testData.rowCount; j++ {
				testCell := board.cells[i][j]

				if testCell.locCol != i || testCell.locRow != j {
					t.Errorf("[%d][%d] is [%d][%d].", i, j, testCell.locCol, testCell.locRow)
				}

				if testCell.hasMine {
					numMines++
				}
			}
		}

		if numMines != testData.mineCount {
			t.Errorf("Incorrect number of mines placed. Is %d, should be %d.", numMines, testData.mineCount)
		}
	}
}
