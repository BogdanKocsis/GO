package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var samples int

	inside := 0
	pi := 0.0
	min := 0.0
	max := 2.0 // consider raza 1
	flag.IntVar(&samples, "samples", 1000000, "Number of samples")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= samples; i++ {
		x := (rand.Float64() * (max - min)) + min
		y := (rand.Float64() * (max - min)) + min
		d := math.Sqrt(math.Pow(x-1, 2) + math.Pow(y-1, 2))
		if d <= 1 {
			inside += 1
		}
		pi = float64(inside) / float64(i) * 4

	}

	fmt.Println(pi)
}
