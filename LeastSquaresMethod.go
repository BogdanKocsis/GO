package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
)

type xy struct {
	x []float64
	y []float64
}

func (d xy) Len() int {
	return len(d.x)
}

func (d xy) XY(i int) (x, y float64) {
	x = d.x[i]
	y = d.y[i]
	return
}

type Point struct {
	x float64
	y float64
}

func LeastSquaresMethod(points *[]Point) (a float64, b float64) {

	n := float64(len(*points))

	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumXX := 0.0

	for _, p := range *points {
		sumX += p.x
		sumY += p.y
		sumXY += p.x * p.y
		sumXX += p.x * p.x
	}

	base := n*sumXX - sumX*sumX
	a = (n*sumXY - sumX*sumY) / base
	b = (sumXX*sumY - sumXY*sumX) / base

	return a, b
}

//func Rotate(p Point, alpha float64) Point {
//	newPoint := Point{
//		x: p.x*math.Cos(alpha) - p.y*math.Sin(alpha),
//		y: p.x*math.Sin(alpha) + p.y*math.Cos(alpha),
//	}
//	return newPoint
//}

func Plot(a float64, b float64, points *[]Point, numberOfPoint int) {

	data := xy{
		x: make([]float64, numberOfPoint),
		y: make([]float64, numberOfPoint),
	}

	var i int
	i = 0
	for _, p := range *points {
		data.x[i] = p.x
		data.y[i] = p.y
		i += 1
	}

	log.Printf("%v*x+%v", a, b)
	line := plotter.NewFunction(func(x float64) float64 { return a*x + b })

	p := plot.New()

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(2)
	scatter, err := plotter.NewScatter(data)
	if err != nil {
		log.Panic(err)
	}
	p.Add(scatter, line)

	w, err := p.WriterTo(300, 300, "svg")
	if err != nil {
		log.Panic(err)
	}

	_, err = w.WriteTo(os.Stdout)
	if err != nil {
		log.Panic(err)
	}

}

func main() {
	points := make([]Point, 0)
	var numberOfPoints int
	//alpha := math.Pi / 2

	file, err := os.Open("points.txt")

	if err != nil {
		fmt.Println(err)
	}

	_, _ = fmt.Fscan(file, &numberOfPoints)

	for i := 0; i < numberOfPoints; i++ {
		var currentPoint Point
		_, _ = fmt.Fscan(file, &currentPoint.x)
		_, _ = fmt.Fscan(file, &currentPoint.y)
		log.Print(currentPoint, " ")
		points = append(points, currentPoint)
	}

	log.Print("\n")
	a, b := LeastSquaresMethod(&points)
	log.Print("Before rotation", "\n")
	log.Println("a= ", a)
	log.Println("b= ", b)
	Plot(a, b, &points, numberOfPoints)

	//for i := 0; i < numberOfPoints; i++ {
	//	points[i] = Rotate(points[i], alpha)
	//}
	//
	//a, b = LeastSquaresMethod(&points)
	//log.Print("After rotation", "\n")
	//log.Println("a= ", a)
	//log.Println("b= ", b)
	//
	//Plot(a, b, &points, numberOfPoints)

}
