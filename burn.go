package termburn

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Run() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetOutputMode(termbox.Output256)
	t := newTerminal()

	ch := make(chan termbox.Event)
	go func() {
		for {
			ch <- termbox.PollEvent()
		}
	}()

loop:
	for {
		select {
		case ev := <-ch:
			switch ev.Type {
			case termbox.EventKey:
				if ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc {
					break loop
				}
			}
			break
		default:
			if err := t.update(); err != nil {
				panic(err)
			}
			time.Sleep(time.Millisecond * 150)
		}
	}
}
