package termburn

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

var colors = []uint16{16, 52, 88, 124, 160, 196, 202, 208, 214, 220, 226, 231}

func Run() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetOutputMode(termbox.Output256)
	colSize, rowSize := termbox.Size()
	rowSize += 2
	cells := make([][]float32, rowSize)
	for i := range cells {
		cells[i] = make([]float32, colSize)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	cells[rowSize-1][0] = rand.Float32()
	for i := 1; i < colSize; i++ {
		val := cells[rowSize-1][i-1] + r.Float32()*0.6 - 0.3
		if val >= 1.0 {
			val = 1.0
		}
		if val <= 0.0 {
			val = 0.0
		}
		cells[rowSize-1][i] = val
	}
	for i := rowSize - 2; i >= 0; i-- {
		for j := 0; j < colSize; j++ {
			num, sum := 1, cells[i+1][j]
			if j-1 > 0 {
				num++
				sum += cells[i+1][j-1]
			}
			if j+1 < colSize-1 {
				num++
				sum += cells[i+1][j+1]
			}
			if i+2 < rowSize-1 {
				num++
				sum += cells[i+1][j]
			}
			val := sum / float32(num)
			val += r.Float32()*0.04 - 0.02
			if val >= 1.0 {
				val = 1.0
			}
			if val <= 0.0 {
				val = 0.0
			}
			cells[i][j] = val
		}
	}

	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			value := cells[i][j]
			color := colors[int(value*float32(len(colors)-1))]
			termbox.SetCell(j, i, ' ', termbox.ColorDefault, termbox.Attribute(color+1))
		}
	}
	if err := termbox.Flush(); err != nil {
		panic(err)
	}

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
		}
	}
}
