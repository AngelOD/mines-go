// Termbox-Go UI for my Minesweeper core (github.com/angelod/mines-go/core)
//
// Heavily inspired by
//   https://github.com/rn2dy/sokoban-go

package main

import (
	"flag"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	animationSpeed = 10 * time.Millisecond
)

var (
	eventQueue chan termbox.Event
)

func main() {
	mineCount := flag.Int("mines", 10, "The number of mines to start with.")

	flag.Parse()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	initEventQueue()
	mainLoop(mineCount)
}

func initEventQueue() {
	eventQueue = make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()
}

func mainLoop(mineCount *int) {
	g := NewGame(30, 20, *mineCount)
	render(g)

	for {
		ev := <-eventQueue
		if ev.Type == termbox.EventKey {
			g.processKeyEvent(ev)
		}

		render(g)

		if g.fsm.Current() == "quitting" {
			return
		}

		time.Sleep(animationSpeed)
	}
}
