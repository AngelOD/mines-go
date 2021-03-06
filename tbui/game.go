package main

import (
	minesCore "github.com/angelod/mines-go/core"
	"github.com/looplab/fsm"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Game struct {
	fsm      *fsm.FSM
	board    *minesCore.Board
	curCol   int
	curRow   int
	debug    bool
	gameWon  bool
	numCols  int
	numRows  int
	numMines int
}

func NewGame(colCount, rowCount, mineCount int) *Game {
	g := &Game{
		numCols:  colCount,
		numRows:  rowCount,
		numMines: mineCount,
		board:    minesCore.NewBoard(colCount, rowCount, mineCount),
		debug:    false,
	}

	g.initFSM()

	return g
}

func (g *Game) checkMove(dx, dy int) {
	var x, y int = g.curCol + dx, g.curRow + dy

	if x < 0 {
		x = 0
	} else if x >= g.numCols {
		x = g.numCols - 1
	}

	if y < 0 {
		y = 0
	} else if y >= g.numRows {
		y = g.numRows - 1
	}

	g.curCol = x
	g.curRow = y
}

func (g *Game) move(dir Direction) {
	if !g.fsm.Is("running") {
		return
	}

	switch dir {
	case UP:
		g.checkMove(0, -1)
	case DOWN:
		g.checkMove(0, 1)
	case LEFT:
		g.checkMove(-1, 0)
	case RIGHT:
		g.checkMove(1, 0)
	}
}

func (g *Game) toggleDebug() bool {
	g.debug = !g.debug
	return g.debug
}

func (g *Game) reveal() {
	_, gameStatus := g.board.Click(g.curCol, g.curRow)

	switch gameStatus {
	case minesCore.GAME_LOST:
		g.fsm.Event("game_lost")

	case minesCore.GAME_WON:
		g.fsm.Event("game_won")
	}
}

func (g *Game) markCell() {
	g.board.MarkCell(g.curCol, g.curRow)
}

func (g *Game) nextLevel() {
	mineCount := g.numMines
	maxMineCount := g.board.GetMaxMineCount()

	if g.fsm.Is("finished") && g.gameWon && mineCount < maxMineCount {
		mineCount++
	}

	g.numMines = mineCount
	g.board = minesCore.NewBoard(g.numCols, g.numRows, g.numMines)
	g.fsm.Event("start_game")
}
