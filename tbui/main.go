// Test of Termbox-Go, heavily inspired by
//   https://github.com/rn2dy/sokoban-go

package main

import (
	"flag"
	"time"

	minesCore "github.com/angelod/mines-go/core"
	"github.com/nsf/termbox-go"
)

const (
	animationSpeed = 10 * time.Millisecond
)

func main() {
	mineCount := flag.Int("mines", 10, "The number of mines to start with.")

	flag.Parse()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

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
