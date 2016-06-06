// Test of Termbox-Go, heavily inspired by
//   https://github.com/rn2dy/sokoban-go

package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

const (
	animationSpeed = 10 * time.Millisecond
)

func main() {
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

	g := NewGame(10, 10, 10)
	render(g)

	for {
		ev := <-eventQueue
		if ev.Type == termbox.EventKey {
			switch {
			case ev.Key == termbox.KeyArrowUp || ev.Ch == 'w':
				g.move(UP)
			case ev.Key == termbox.KeyArrowDown || ev.Ch == 's':
				g.move(DOWN)
			case ev.Key == termbox.KeyArrowLeft || ev.Ch == 'a':
				g.move(LEFT)
			case ev.Key == termbox.KeyArrowRight || ev.Ch == 'd':
				g.move(RIGHT)
			case ev.Key == termbox.KeySpace:
				g.reveal()
			case ev.Key == termbox.KeyEsc:
				return
			}
		}

		render(g)
		time.Sleep(animationSpeed)
	}
}
