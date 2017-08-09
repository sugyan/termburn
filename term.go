package termburn

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

var colors = []uint16{16, 52, 88, 124, 160, 196, 202, 208, 214, 220, 226, 231}

type terminal struct {
	col, row int
	values   [][]float32
}

func newTerminal() *terminal {
	col, row := termbox.Size()
	row++
	values := make([][]float32, row)
	for i := 0; i < row; i++ {
		values[i] = make([]float32, col)
	}
	for i := 0; i < col; i++ {
		values[row-1][i] = rand.Float32()
	}
	return &terminal{
		col:    col,
		row:    row,
		values: values,
	}
}

func (t *terminal) update() (err error) {
	for i := 0; i < t.col; i++ {
		val := t.values[t.row-1][i] + rand.Float32()*0.5 - 0.25
		t.values[t.row-1][i] = clip(val)
	}
	for i := t.row - 2; i >= 0; i-- {
		for j := 0; j < t.col; j++ {
			num, sum := 1, t.values[i+1][j]
			if j-1 > 0 {
				num++
				sum += t.values[i+1][j-1]
			}
			if j+1 < t.col-1 {
				num++
				sum += t.values[i+1][j+1]
			}
			if i+2 < t.row-1 {
				num++
				sum += t.values[i+1][j]
			}
			val := sum / float32(num)
			val += 0.1*(rand.Float32()-0.5) - 0.5/float32(t.row)
			t.values[i][j] = clip(val)
		}
	}
	return t.render()
}

func (t *terminal) render() (err error) {
	for i := 0; i < t.row-1; i++ {
		for j := 0; j < t.col; j++ {
			value := t.values[i][j]
			color := colors[int(value*float32(len(colors)-1))]
			termbox.SetCell(j, i, ' ', termbox.ColorDefault, termbox.Attribute(color+1))
		}
	}
	if err = termbox.Flush(); err != nil {
		return
	}
	return nil
}

func clip(value float32) float32 {
	if value > 1.0 {
		return 1.0
	}
	if value < 0.0 {
		return 0.0
	}
	return value
}
