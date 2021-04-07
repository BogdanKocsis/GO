package main

import (
	r "./Interpolare"
	"bufio"
	"fmt"
	"math"
	"os"
)

type Points struct {
	x, y, z float64
}

func calculateDistance(firstPoint Points, secondPoint Points) float64 {
	return math.Sqrt(math.Pow(firstPoint.x-secondPoint.x, 2) + math.Pow(firstPoint.y-secondPoint.y, 2) + math.Pow(firstPoint.z-secondPoint.z, 2))
}

func calculateArea(a, b, c float64) float64 {
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func main() {

	p := 20

	var nrPuncteDeasupra int
	var nrPuncteDedesubt int
	var nrFelii int

	var currentX float64
	var currentY float64

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fscan(file, &nrFelii)

	outputFile, _ := os.Create("result.obj")
	w := bufio.NewWriter(outputFile)
	var vertices []Points

	for ii := 0; ii < nrFelii; ii++ {

		var x1 []float64
		var y1 []float64
		var x2 []float64
		var y2 []float64

		fmt.Fscan(file, &nrPuncteDeasupra)
		fmt.Fscan(file, &nrPuncteDedesubt)

		for i := 0; i < nrPuncteDeasupra; i++ {
			fmt.Fscan(file, &currentX)
			fmt.Fscan(file, &currentY)
			x1 = append(x1, currentX)
			y1 = append(y1, currentY)
		}

		for i := 0; i < nrPuncteDedesubt; i++ {
			fmt.Fscan(file, &currentX)
			fmt.Fscan(file, &currentY)
			x2 = append(x2, currentX)
			y2 = append(y2, currentY)
		}

		var s1 = r.NewSpline(x1, y1, r.CubicSecondDeriv, 0, 0)
		var s2 = r.NewSpline(x2, y2, r.CubicSecondDeriv, 0, 0)

		//fmt.Println(s.At(7))

		var h = (x1[nrPuncteDeasupra-1] - x1[0]) / float64(p)

		for i := 1; i < p; i++ {
			xTemp := x1[0] + float64(i)*h
			yTemp := s1.At(xTemp)
			//fmt.Print("v ", xTemp, " ", yTemp, " ", 100, "\n")
			fmt.Fprintf(w, "v %f %f %d \n", xTemp, yTemp, ii*4)
			vertices = append(vertices, Points{xTemp, yTemp, float64(ii * 4)})
		}

		h = (x2[nrPuncteDedesubt-1] - x2[0]) / float64(p)

		for i := 1; i < p; i++ {
			xTemp := x2[0] + float64(i)*h
			yTemp := s2.At(xTemp)
			//fmt.Print("v ", xTemp, " ", yTemp, " ", 100, "\n")
			fmt.Fprintf(w, "v %f %f %d \n", xTemp, yTemp, ii*4)
			vertices = append(vertices, Points{xTemp, yTemp, float64(ii * 4)})
		}
	}

	area := .0
	p--
	for i := 1; i < p; i++ {

		//fmt.Fprintf(w, "f %d %d %d \n", i, 2*p+i, 2*p+i+1 )
		//fmt.Fprintf(w, "f %d %d %d \n", i, i+1, 2*p+i+1 )
		//fmt.Fprintf(w, "f %d %d %d \n", p+i, 3*p+i, 3*p+i+1 )
		//fmt.Fprintf(w, "f %d %d %d \n", p+i, p+i+1, 3*p+i+1 )

		fmt.Fprintf(w, "f %d %d %d \n", i, 2*p+i, 2*p+i+1)
		a := calculateDistance(vertices[i-1], vertices[2*p+i-1])
		b := calculateDistance(vertices[2*p+i-1], vertices[2*p+i])
		c := calculateDistance(vertices[2*p+i], vertices[i-1])
		area += calculateArea(a, b, c)

		fmt.Fprintf(w, "f %d %d %d \n", i, i+1, 2*p+i+1)
		a = calculateDistance(vertices[i-1], vertices[i])
		b = calculateDistance(vertices[i], vertices[2*p+i])
		area += calculateArea(a, b, c)

		fmt.Fprintf(w, "f %d %d %d \n", p+i, 3*p+i, 3*p+i+1)
		a = calculateDistance(vertices[p+i-1], vertices[3*p+i-1])
		b = calculateDistance(vertices[3*p+i-1], vertices[3*p+i])
		c = calculateDistance(vertices[3*p+i], vertices[p+i-1])
		area += calculateArea(a, b, c)

		fmt.Fprintf(w, "f %d %d %d \n", p+i, p+i+1, 3*p+i+1)

		a = calculateDistance(vertices[p+i-1], vertices[p+i])
		b = calculateDistance(vertices[p+i], vertices[3*p+i])
		area += calculateArea(a, b, c)
	}
	p++
	fmt.Fprintf(w, "f %d %d %d \n", 1, p, 3*p-2)
	a := calculateDistance(vertices[0], vertices[p-1])
	b := calculateDistance(vertices[p-1], vertices[3*p-3])
	c := calculateDistance(vertices[3*p-3], vertices[0])
	area += calculateArea(a, b, c)

	fmt.Fprintf(w, "f %d %d %d \n", 1, 2*p-1, 3*p-2)
	a = calculateDistance(vertices[0], vertices[2*p-2])
	b = calculateDistance(vertices[2*p-2], vertices[3*p-3])
	area += calculateArea(a, b, c)

	fmt.Fprintf(w, "f %d %d %d \n", p-1, 3*p-3, 4*p-4)
	a = calculateDistance(vertices[p-2], vertices[3*p-4])
	b = calculateDistance(vertices[3*p-4], vertices[4*p-5])
	c = calculateDistance(vertices[4*p-5], vertices[p-2])
	area += calculateArea(a, b, c)

	fmt.Fprintf(w, "f %d %d %d \n", p-1, 2*p-2, 4*p-4)
	a = calculateDistance(vertices[p-2], vertices[2*p-3])
	b = calculateDistance(vertices[2*p-3], vertices[4*p-5])
	area += calculateArea(a, b, c)

	fmt.Println("Calculated Area: ", area)
	w.Flush()

}
