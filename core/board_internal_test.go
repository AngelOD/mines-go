package core

import (
	"testing"
)

type cellCheckEntry struct {
	checkCol         int
	checkRow         int
	expectedCellType CellType
	expectedCellNum  int
}

var boardFindCellTypeTests = []struct {
	revealFirst      bool
	col              int
	row              int
	expectedCellType CellType
	expectedCellNum  int
}{
	{true, 0, 0, HAS_MINE, -1},
	{true, 1, 1, HAS_NUMBER, 8},
	{true, 0, 9, HAS_NUMBER, 0},
	{true, 9, 9, HAS_NUMBER, 0},
	{true, 9, 7, HAS_NUMBER, 1},
	{true, 9, 5, HAS_MINE, -1},
	{false, 3, 3, NOT_REVEALED, -1},
	{false, 9, 0, NOT_REVEALED, -1},
}

// 0-8: Expected number
// 9: Bomb
var boardTestLayout = [][]int{
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
	{10, 10, 9, 9, 99},
	{10, 20, 0, 0, 0},
	{10, 20, 1, 0, 1},
	{10, 20, 0, 1, 10},
	{10, 20, 1, 1, 11},
	{10, 20, 9, 19, 199},
	{20, 10, 0, 0, 0},
	{20, 10, 1, 0, 1},
	{20, 10, 0, 1, 20},
	{20, 10, 1, 1, 21},
	{20, 10, 19, 9, 199},
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
	{10, 10, 99, 9, 9},
	{10, 20, 0, 0, 0},
	{10, 20, 1, 1, 0},
	{10, 20, 10, 0, 1},
	{10, 20, 11, 1, 1},
	{10, 20, 199, 9, 19},
	{20, 10, 0, 0, 0},
	{20, 10, 1, 1, 0},
	{20, 10, 20, 0, 1},
	{20, 10, 21, 1, 1},
	{20, 10, 199, 19, 9},
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

var boardRevealCellTests = []struct {
	revealCol  int
	revealRow  int
	checkPairs []cellCheckEntry
}{
	{0, 0, []cellCheckEntry{
		{0, 0, HAS_MINE, -1},
		{1, 0, NOT_REVEALED, -1},
		{0, 1, NOT_REVEALED, -1},
		{1, 1, NOT_REVEALED, -1},
	}},
	{0, 9, []cellCheckEntry{
		{0, 9, HAS_NUMBER, 0},
		{0, 8, HAS_NUMBER, 0},
		{0, 7, HAS_NUMBER, 2},
		{0, 6, NOT_REVEALED, -1},
		{1, 9, HAS_NUMBER, 0},
		{1, 8, HAS_NUMBER, 0},
		{1, 7, HAS_NUMBER, 3},
		{1, 6, NOT_REVEALED, -1},
		{2, 9, HAS_NUMBER, 0},
		{2, 8, HAS_NUMBER, 0},
		{2, 7, HAS_NUMBER, 2},
		{2, 6, NOT_REVEALED, -1},
		{3, 9, HAS_NUMBER, 1},
		{3, 8, HAS_NUMBER, 1},
		{3, 7, NOT_REVEALED, -1},
		{4, 9, NOT_REVEALED, -1},
		{4, 8, NOT_REVEALED, -1},
		{4, 7, NOT_REVEALED, -1},
	}},
}

func TestBoardFindCellType(t *testing.T) {
	var board *Board

	board = setupTestBoard()

	for _, testData := range boardFindCellTypeTests {
		if testData.revealFirst {
			board.revealCell(testData.col, testData.row)
		}

		board.findCellType(testData.col, testData.row)

		actualCellType := board.cells[board.GetIndexFromLocation(testData.col, testData.row)].CType
		actualCellNum := board.cells[board.GetIndexFromLocation(testData.col, testData.row)].CNum

		if actualCellNum != testData.expectedCellNum || actualCellType != testData.expectedCellType {
			t.Errorf("[%d][%d] Expected (%d)(%d). Got (%d)(%d).",
				testData.col, testData.row,
				testData.expectedCellType, testData.expectedCellNum,
				actualCellType, actualCellNum)
		}
	}
}

func TestBoardFindProximityMines(t *testing.T) {
	var board *Board

	board = setupTestBoard()

	// Run actual test
	for i := 0; i < board.colCount; i++ {
		for j := 0; j < board.rowCount; j++ {
			expectedMineCount := boardTestLayout[j][i]

			if expectedMineCount < 9 {
				board.findProximityMineCount(i, j)
				actualMineCount := board.cells[board.GetIndexFromLocation(i, j)].proximityMineCount

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
		index := board.GetIndexFromLocation(testData.col, testData.row)

		if index != testData.expectedIndex {
			t.Errorf("[%d][%d] Expected index %d. Got %d.", testData.col, testData.row, testData.expectedIndex, index)
		}
	}
}

func TestBoardGetLocationFromIndex(t *testing.T) {
	for _, testData := range boardGetLocationFromIndexTests {
		board := NewBoard(testData.colCount, testData.rowCount, 10)
		col, row := board.GetLocationFromIndex(testData.index)

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
				testCell := board.cells[board.GetIndexFromLocation(i, j)]

				if testCell.Col != i || testCell.Row != j {
					t.Errorf("[%d][%d] is [%d][%d].", i, j, testCell.Col, testCell.Row)
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

func TestBoardRevealCell(t *testing.T) {
	for _, testData := range boardRevealCellTests {
		board := setupTestBoard()
		board.revealCell(testData.revealCol, testData.revealRow)

		for _, cellData := range testData.checkPairs {
			board.findCellType(cellData.checkCol, cellData.checkRow)

			testCell := &board.cells[board.GetIndexFromLocation(cellData.checkCol, cellData.checkRow)]
			actualCellType := testCell.CType
			actualCellNum := testCell.CNum

			if actualCellType != cellData.expectedCellType || actualCellNum != cellData.expectedCellNum {
				t.Errorf("[%d][%d] Expected (%d)(%d). Got (%d)(%d).",
					cellData.checkCol, cellData.checkRow,
					cellData.expectedCellType, cellData.expectedCellNum,
					actualCellType, actualCellNum)
			}
		}
	}
}

// ====================== HELPER METHODS ======================

func setupTestBoard() (board *Board) {
	var colCount, rowCount int = len(boardTestLayout), len(boardTestLayout[0])
	board = NewBoard(colCount, rowCount, 1)

	// Reset board to our desired layout
	for i := 0; i < colCount; i++ {
		for j := 0; j < rowCount; j++ {
			hasMine := false
			if boardTestLayout[j][i] == 9 {
				hasMine = true
			}

			board.cells[board.GetIndexFromLocation(i, j)].hasMine = hasMine
		}
	}

	return
}
