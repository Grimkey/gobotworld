package main

import (
	"fmt"
	"gobotworld/src/terminal"
	"gobotworld/src/world"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func panicOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func main() {
	// Create a file
	file, err := os.Create("game.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)

	gameWorld := world.DefaultWorld(logger)

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
					gameWorld.Move(gameWorld.Player, world.West)
				case tcell.KeyRight:
					gameWorld.Move(gameWorld.Player, world.East)
				case tcell.KeyUp:
					gameWorld.Move(gameWorld.Player, world.North)
				case tcell.KeyDown:
					gameWorld.Move(gameWorld.Player, world.South)

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
		gameWorld.Tick()
		gameWorld.NpcMove()
		term.DrawWorld(gameWorld)
		term.Show()
		cnt++
		dur += time.Since(start)
	}
}
