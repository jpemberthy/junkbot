package main

import (
	"math"

	"github.com/gizak/termui"
)

func main() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	p := termui.NewPar(":PRESS q TO QUIT DEMO")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = termui.ColorWhite
	p.BorderLabel = "Text Box"
	p.BorderFg = termui.ColorCyan
	p.Handle("/timer/1s", func(e termui.Event) {
		cnt := e.Data.(termui.EvtTimer)
		if cnt.Count%2 == 0 {
			p.TextFgColor = termui.ColorRed
		} else {
			p.TextFgColor = termui.ColorWhite
		}
	})

	sinps := (func() []float64 {
		n := 220
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	})()

	lc := termui.NewLineChart()
	lc.BorderLabel = "dot-mode Line Chart"
	lc.Data = sinps
	lc.Width = 50
	lc.Height = 11
	lc.X = 0
	lc.Y = 14
	lc.AxesColor = termui.ColorWhite
	lc.LineColor = termui.ColorRed | termui.AttrBold
	lc.Mode = "dot"

	bc := termui.NewBarChart()
	bcdata := []int{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bclabels := []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.BorderLabel = "Bar Chart"
	bc.Width = 26
	bc.Height = 10
	bc.X = 51
	bc.Y = 0
	bc.DataLabels = bclabels
	bc.BarColor = termui.ColorGreen
	bc.NumColor = termui.ColorBlack

	lc1 := termui.NewLineChart()
	lc1.BorderLabel = "braille-mode Line Chart"
	lc1.Data = sinps
	lc1.Width = 26
	lc1.Height = 11
	lc1.X = 51
	lc1.Y = 14
	lc1.AxesColor = termui.ColorWhite
	lc1.LineColor = termui.ColorYellow | termui.AttrBold

	draw := func(t int) {
		lc.Data = sinps[t/2%220:]
		lc1.Data = sinps[2*t%220:]
		bc.Data = bcdata[t/2%10:]
		termui.Render(p, bc, lc1, lc)
	}
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Handle("/timer/1s", func(e termui.Event) {
		t := e.Data.(termui.EvtTimer)
		draw(int(t.Count))
	})
	termui.Loop()
}
