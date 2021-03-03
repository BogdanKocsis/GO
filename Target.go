package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var samples int
	//var verbose bool
	inside := 0
	pi := 0.0

	flag.IntVar(&samples, "samples", 1000000, "Number of samples")
	//flag.BoolVar(&verbose, "v", false, "Verbose display")

	flag.Parse()

	for i := 1; i <= samples; i++ {
		var x = rand.Float64()
		var y = rand.Float64()
		d := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
		if d <= 1 {
			inside += 1
		}
		pi = float64(inside) / float64(i) * 4
		//if verbose {
		//	fmt.Printf("Occurrence number: %d, pi estimation : %f\n", i, pi)
		//}
	}
	fmt.Println(pi)
}
