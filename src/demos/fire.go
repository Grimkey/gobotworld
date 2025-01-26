package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"gobotworld/src/terminal"
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
	fmt.Println("fire demo")

	term, err := terminal.Init()
	defer term.Fini()
	panicOnError(err)

	fire1 := tcell.NewRGBColor(0xFA, 0xC0, 0x00)
	if !fire1.IsRGB() {
		panic("oops")
	}
	fire1Style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(fire1)

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
				}
			case *tcell.EventResize:
				term.Show()
			}
		}
	}()

	term.SetCell(1, 1, fire1Style, '^')

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}
		term.Show()
	}
}
