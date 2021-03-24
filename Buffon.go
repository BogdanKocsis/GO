package main

import (
	. "fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math"
	"math/rand"
)

func main() {
	var attempts int
	var l float64
	var events int
	Println("Input attempts number:")
	Scanf("%d\n", &attempts)
	Println("Input needle len:")
	Scanf("%f\n", &l)
	rand.Seed(rand.Int63())
	pts := make(plotter.XYs, attempts)
	for i := 1; i < attempts; i++ {
		pts[i].X = float64(i)
		x := rand.Float64() * l
		radian := float64(rand.Intn(180)) * (math.Pi / 180.0)
		y := x + l*math.Sin(radian)
		if int(x/l) < int(y/l) {
			events++
		}
		if events != 0 {
			pi := (2 * float64(i)) / (float64(events))
			pts[i].Y = pi
		}

	}

	pi := (2 * float64(attempts)) / (float64(events))
	Println(pi)
	var (
		p = plot.New()
	)

	p.Title.Text = "Buffon problem"
	p.X.Label.Text = "attempt"
	p.Y.Label.Text = "Pi approximation"

	_ = plotutil.AddLinePoints(p,
		"Pi on attempt", pts)
	p.Y.Min = 2
	p.Y.Max = 4
	p.Y.Dashes = []vg.Length{}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "buffon.png"); err != nil {
		panic(err)

	}
}
