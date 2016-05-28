package core

import (
	"testing"
)

// 0-8: Expected number
// 9: Bomb
var boardFindProximityMineCountTestLayout = [][]int{
	{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
	{9, 8, 9, 7, 6, 5, 6, 9, 9, 4},
	{9, 9, 9, 9, 9, 9, 3, 9, 9, 2},
	{3, 5, 4, 5, 4, 4, 3, 4, 3, 2},
	{9, 3, 9, 3, 9, 3, 9, 3, 9, 2},
	{4, 9, 5, 9, 3, 9, 5, 9, 5, 9},
	{9, 9, 9, 2, 2, 2, 9, 9, 9, 2},
	{2, 3, 2, 2, 2, 3, 3, 3, 2, 1},
	{0, 0, 0, 1, 9, 9, 1, 0, 0, 0},
	{0, 0, 0, 1, 2, 2, 1, 0, 0, 0},
}

var boardFindSurroundingCellsTests = []struct {
	colCount      int
	rowCount      int
	col           int
	row           int
	expectedCount int
}{
	{10, 10, 0, 0, 3},
	{10, 10, 1, 0, 5},
	{10, 10, 0, 1, 5},
	{10, 10, 1, 1, 8},
	{10, 10, 9, 9, 3},
	{10, 10, 8, 9, 5},
	{10, 10, 9, 8, 5},
	{10, 10, 8, 8, 8},
}

var boardGetIndexFromLocationTests = []struct {
	colCount      int
	rowCount      int
	col           int
	row           int
	expectedIndex int
}{
	{10, 10, 0, 0, 0},
	{10, 10, 1, 0, 1},
	{10, 10, 0, 1, 10},
	{10, 10, 1, 1, 11},
}

var boardGetLocationFromIndexTests = []struct {
	colCount    int
	rowCount    int
	index       int
	expectedCol int
	expectedRow int
}{
	{10, 10, 0, 0, 0},
	{10, 10, 1, 1, 0},
	{10, 10, 10, 0, 1},
	{10, 10, 11, 1, 1},
}

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

var boardIsValidCoordTests = []struct {
	colCount int
	rowCount int
	col      int
	row      int
	expected bool
}{
	{10, 10, 0, 0, true},
	{10, 10, 5, 5, true},
	{10, 10, 9, 9, true},
	{10, 10, -1, 0, false},
	{10, 10, 0, -1, false},
	{10, 10, -1, -1, false},
	{10, 10, 10, 9, false},
	{10, 10, 9, 10, false},
	{10, 10, 10, 10, false},
}

func TestBoardFindProximityMines(t *testing.T) {
	var colCount, rowCount int = 10, 10

	board := NewBoard(colCount, rowCount, 1)

	// Reset board so we have full control over it
	for i := 0; i < colCount; i++ {
		for j := 0; j < rowCount; j++ {
			hasMine := false
			if boardFindProximityMineCountTestLayout[j][i] == 9 {
				hasMine = true
			}

			board.cells[board.getIndexFromLocation(i, j)].hasMine = hasMine
		}
	}

	// Run actual test
	for i := 0; i < colCount; i++ {
		for j := 0; j < rowCount; j++ {
			expectedMineCount := boardFindProximityMineCountTestLayout[j][i]

			if expectedMineCount < 9 {
				board.findProximityMineCount(i, j)
				actualMineCount := board.cells[board.getIndexFromLocation(i, j)].proximityMineCount

				if actualMineCount != expectedMineCount {
					t.Errorf("[%d][%d] Expected %d, found %d.", i, j, expectedMineCount, actualMineCount)
				}
			}
		}
	}
}

func TestBoardFindSurroundingCells(t *testing.T) {
	for _, testData := range boardFindSurroundingCellsTests {
		board := NewBoard(testData.colCount, testData.rowCount, 10)
		cells := board.findSurroundingCells(testData.col, testData.row)

		if len(cells) != testData.expectedCount {
			t.Errorf("Expected %d cells. Found %d.", testData.expectedCount, len(cells))
		}
	}
}

func TestBoardGetIndexFromLocation(t *testing.T) {
	for _, testData := range boardGetIndexFromLocationTests {
		board := NewBoard(testData.colCount, testData.rowCount, 10)
		index := board.getIndexFromLocation(testData.col, testData.row)

		if index != testData.expectedIndex {
			t.Errorf("[%d][%d] Expected index %d. Got %d.", testData.col, testData.row, testData.expectedIndex, index)
		}
	}
}

func TestBoardGetLocationFromIndex(t *testing.T) {
	for _, testData := range boardGetLocationFromIndexTests {
		board := NewBoard(testData.colCount, testData.rowCount, 10)
		col, row := board.getLocationFromIndex(testData.index)

		if col != testData.expectedCol || row != testData.expectedRow {
			t.Errorf("[%d] Expected [%d][%d]. Got [%d][%d].", testData.index, testData.expectedCol, testData.expectedRow, col, row)
		}
	}
}

func TestBoardInit(t *testing.T) {
	for _, testData := range boardInitTests {
		var numMines int

		board := NewBoard(testData.colCount, testData.rowCount, testData.mineCount)

		for i := 0; i < testData.colCount; i++ {
			for j := 0; j < testData.rowCount; j++ {
				testCell := board.cells[board.getIndexFromLocation(i, j)]

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

func TestBoardIsValidCoord(t *testing.T) {
	for _, testData := range boardIsValidCoordTests {
		board := NewBoard(testData.colCount, testData.rowCount, 1)
		result := board.isValidCoord(testData.col, testData.row)

		if result != testData.expected {
			t.Errorf("Expected [%d][%d] to be %t, but it wasn't.", testData.col, testData.row, testData.expected)
		}
	}
}
