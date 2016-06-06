package core

type CellType int

const (
	NOT_REVEALED CellType = -1
	HAS_NUMBER   CellType = 0
	HAS_MINE     CellType = 9
)

type GameStatus int

const (
	GAME_LOST    GameStatus = -1
	GAME_RUNNING GameStatus = 0
	GAME_WON     GameStatus = 1
)
