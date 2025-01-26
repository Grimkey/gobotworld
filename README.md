# Bot World game

This is just a fun side project. It is not meant to be used by anyone except the author. The purpose of this is to just program and have fun. It is meant to be a very low dependency game that I just add new features to whenever I need to code.

## Running

To start the game just run it in code, I usually run it like:

```
go run src/main.go
```

This will run it without creating an executable.

### Navigating

The arrow keys allow you to move your character around.

### Exiting the game

You can use either the "enter" or "escape" key to exit out of the game. 


### Key mapping

You can find the key mapping code in main.go if you want to update it. Here is an example of how it is setup. The `quit` variable is a channel that is monitored by the main loop. `gameWorld.Move` will queue up a change to be handled the next loops through.

```go
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
```

# Understanding the main loop

The main loop code is in `src/main.go`. 

```go
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
```

The main loop is on a 50 millisecond timer this controls the loop speed, you can make it faster, but this worked for the testing that I've been doing.

gameWorld.Tick is meant to be a timer within the game itself. I've been playing with day and night behavior. Tick lets us count how many 50ms cycles have passed since the start of the game.


