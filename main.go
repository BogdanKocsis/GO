package main

import (
	. "./Interpolare"
	"bufio"
	"fmt"
	"os"
)

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

		var s1 = NewSpline(x1, y1, CubicSecondDeriv, 0, 0)
		var s2 = NewSpline(x2, y2, CubicSecondDeriv, 0, 0)

		//fmt.Println(s.At(7))

		var h = (x1[nrPuncteDeasupra-1] - x1[0]) / float64(p)

		for i := 1; i < p; i++ {
			xTemp := x1[0] + float64(i)*h
			yTemp := s1.At(xTemp)
			//fmt.Print("v ", xTemp, " ", yTemp, " ", 100, "\n")
			fmt.Fprintf(w, "v %f %f %d \n", xTemp, yTemp, ii*100)
		}

		h = (x2[nrPuncteDedesubt-1] - x2[0]) / float64(p)

		for i := 1; i < p; i++ {
			xTemp := x2[0] + float64(i)*h
			yTemp := s2.At(xTemp)
			//fmt.Print("v ", xTemp, " ", yTemp, " ", 100, "\n")
			fmt.Fprintf(w, "v %f %f %d \n", xTemp, yTemp, ii*100)
		}
	}
	w.Flush()
}
