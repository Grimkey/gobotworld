package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"gobotworld/src/terminal"
	"gobotworld/src/world"
	"gobotworld/src/world/character"
	"os"
	"time"
)

func panicOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func main() {
	gameWorld := world.DefaultWorld()

	term, err := terminal.Init()
	defer term.Fini()

	panicOnError(err)

	quit := make(chan struct{})
	go func() {
		for {
			ev := term.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyLeft:
					gameWorld.Move(character.Player, -1, 0)
				case tcell.KeyRight:
					gameWorld.Move(character.Player, 1, 0)
				case tcell.KeyUp:
					gameWorld.Move(character.Player, 0, -1)
				case tcell.KeyDown:
					gameWorld.Move(character.Player, 0, 1)

				}
			case *tcell.EventResize:
				term.Show()
			}
		}
	}()

	cnt := 0
	dur := time.Duration(0)

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}
		start := time.Now()
		term.DrawWorld(gameWorld)
		term.Show()
		cnt++
		dur += time.Now().Sub(start)
	}
}
