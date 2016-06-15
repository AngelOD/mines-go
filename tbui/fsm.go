package main

import (
	"github.com/looplab/fsm"
	"github.com/nsf/termbox-go"
)

func (g *Game) initFSM() {
	g.fsm = fsm.NewFSM(
		"running",
		fsm.Events{
			{Name: "game_lost", Src: []string{"running"}, Dst: "finished"},
			{Name: "game_won", Src: []string{"running"}, Dst: "finished"},
			{Name: "start_game", Src: []string{"finished"}, Dst: "running"},
			{Name: "quit", Src: []string{"running", "finished"}, Dst: "quitting"},
		},
		fsm.Callbacks{
			"after_game_won": func(e *fsm.Event) {
				g.gameWon = true
			},
			"enter_running": func(e *fsm.Event) {
				g.gameWon = false
			},
		},
	)
}

func (g *Game) processKeyEvent(ev termbox.Event) {
	curState := g.fsm.Current()

	if curState == "running" {
		switch {
		case ev.Key == termbox.KeyArrowUp || ev.Ch == 'w':
			g.move(UP)
		case ev.Key == termbox.KeyArrowDown || ev.Ch == 's':
			g.move(DOWN)
		case ev.Key == termbox.KeyArrowLeft || ev.Ch == 'a':
			g.move(LEFT)
		case ev.Key == termbox.KeyArrowRight || ev.Ch == 'd':
			g.move(RIGHT)
		case ev.Ch == 'c':
			g.toggleDebug()
		case ev.Key == termbox.KeySpace:
			g.reveal()
		case ev.Key == termbox.KeyEsc:
			return
		}
	} else if curState == "finished" {
		switch {
		case ev.Ch == 'c':
			g.toggleDebug()
		case ev.Key == termbox.KeySpace:
			g.nextLevel()
		case ev.Key == termbox.KeyEsc:
			return
		}
	}
}
