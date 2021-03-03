package main

import (
	. "fmt"
	"math"
	"math/rand"
)

func main() {
	var attempts int
	var lineSpacing, l float64
	var events int
	Println("Input attempts number:")
	Scanf("%d\n", &attempts)
	Println("Input Line Spacing len:")
	Scanf("%f\n", &lineSpacing)
	Println("Input needle len:")
	Scanf("%f\n", &l)
	rand.Seed(rand.Int63())

	for i := 1; i < attempts; i++ {

		x := rand.Float64() * lineSpacing / 2
		radian := float64(rand.Intn(180)) * (math.Pi / 180.0)
		y := l / 2 * math.Sin(radian)
		if x <= y {
			events++
		}

	}

	pi := (2 * l * float64(attempts)) / (float64(events) * lineSpacing)
	Println(pi)
}
