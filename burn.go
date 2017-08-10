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

	ch := make(chan termbox.Event)
	go func() {
		for {
			ch <- termbox.PollEvent()
		}
	}()

	width, height := termbox.Size()
	t := newTerminal(width, height)
loop:
	for {
		select {
		case ev := <-ch:
			switch ev.Type {
			case termbox.EventKey:
				if ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc {
					break loop
				}
			case termbox.EventResize:
				t = newTerminal(ev.Width, ev.Height)
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
