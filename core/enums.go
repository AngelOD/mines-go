package core

type CellType int

const (
	NOT_REVEALED CellType = -1
	HAS_NUMBER   CellType = 0
	HAS_MINE     CellType = 9
)
