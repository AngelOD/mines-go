package main

import (
	"bytes"
	"fmt"
	"strconv"

	minesCore "github.com/angelod/mines-go/core"
	"github.com/nsf/termbox-go"
)

var title = "-- Mines-Go (%dx%d) [%d] --"

const (
	debugConsoleColor = termbox.ColorBlack
	debugTextColor    = termbox.ColorWhite
	textColor         = termbox.ColorWhite
	backgroundColor   = termbox.ColorBlue
	viewStartX        = 1
	viewStartY        = 1
	titleStartX       = viewStartX
	titleStartY       = viewStartY
	boardStartX       = viewStartX
	boardStartY       = titleStartY + 2
	instructionStartX = boardStartX
	instructionStartY = boardStartY
)

var boxDrawing = map[string]rune{
	"h":  '\u2500',
	"v":  '\u2502',
	"ul": '\u250C',
	"ur": '\u2510',
	"dl": '\u2514',
	"dr": '\u2518',
}

var tokenColor = map[byte]termbox.Attribute{
	'u': termbox.ColorWhite,
	'm': termbox.ColorRed,
	'n': termbox.ColorGreen,
	'p': termbox.ColorMagenta,
}

var instructions = []string{
	"Instructions:",
	"→ or d    :move right",
	"← or a    :move left",
	"↑ or w    :move up",
	"↓ or s    :move down",
	"space     :reveal",
	"x         :(un)mark cell",
	"     c    :show debug console",
	"     esc  :quit",
	"",
	"Avoid the mines.. That's all.",
}

func render(g *Game) {
	var x, y, dx, dy int
	var cellToken byte
	var cellRune rune

	termbox.Clear(backgroundColor, backgroundColor)

	printText(titleStartX, titleStartY, textColor, backgroundColor, fmt.Sprintf(title, g.numCols, g.numRows, g.numMines))

	x = boardStartX
	y = boardStartY
	board := g.board.GetBoard()
	for _, cell := range board {
		dx = x + cell.Col
		dy = y + cell.Row
		cellToken = 'u'
		cellRune = ' '

		switch cell.CType {
		case minesCore.HAS_MINE:
			cellToken = 'm'
			cellRune = 'X'
		case minesCore.HAS_NUMBER:
			cellToken = 'n'
			if cell.CNum > 0 {
				cellRune = rune(strconv.Itoa(cell.CNum)[0])
			}
		case minesCore.IS_MARKED:
			cellToken = 'm'
			cellRune = '?'
		}

		if cell.Col == g.curCol && cell.Row == g.curRow {
			cellToken = 'p'
		}

		termbox.SetCell(dx, dy, cellRune, textColor, tokenColor[cellToken])
	}

	x = instructionStartX + g.numCols + 4
	y = instructionStartY
	if g.fsm.Is("running") {
		for i, msg := range instructions {
			printText(x, y+i, textColor, backgroundColor, msg)
		}
	} else {
		textToPrint := ""
		textFg, textBg := textColor, backgroundColor

		if g.fsm.Is("finished") && !g.gameWon {
			textToPrint = "GAME OVER!"
			textFg, textBg = termbox.ColorWhite, termbox.ColorRed
		} else {
			textToPrint = "YOU WON!"
			textFg, textBg = termbox.ColorBlack, termbox.ColorGreen
		}
		printTextBox(x, y, textColor, backgroundColor, textFg, textBg, textToPrint)
	}

	termbox.Flush()
}

func printText(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printBox(x, y, w, h int, fg, bg termbox.Attribute) {
	var startRune, endRune, fillerRune rune
	var buf bytes.Buffer
	ulx, uly, drx, dry := x, y, x+w-1, y+h-1

	for curLine := uly; curLine <= dry; curLine++ {
		buf.Reset()

		if curLine == uly {
			startRune = boxDrawing["ul"]
			endRune = boxDrawing["ur"]
			fillerRune = boxDrawing["h"]
		} else if curLine == dry {
			startRune = boxDrawing["dl"]
			endRune = boxDrawing["dr"]
			fillerRune = boxDrawing["h"]
		} else {
			startRune = boxDrawing["v"]
			endRune = boxDrawing["v"]
			fillerRune = ' '
		}

		buf.WriteRune(startRune)
		for i := ulx + 1; i < drx; i++ {
			buf.WriteRune(fillerRune)
		}
		buf.WriteRune(endRune)

		printText(ulx, curLine, fg, bg, buf.String())
	}
}

func printTextBox(x, y int, boxFg, boxBg, textFg, textBg termbox.Attribute, msg string) {
	width, height := len(msg)+4, 5
	textX, textY := x+2, y+2

	printBox(x, y, width, height, boxFg, boxBg)
	printText(textX, textY, textFg|termbox.AttrBold, textBg, msg)
}
