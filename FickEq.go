package main

//This code is for susceptible/infected population.
//The infected may disperse in 1D via Fick's law.
//Newton's method is used.
//The full Jacobian matrix is defined.
//The linear steps are solved by A\d.

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {

	var sus0 = 60.0
	var inf0 = 0
	a := 20.0 / 50.0
	b := 1.0
	var D = 1000.0
	var n = 200
	var nn = 2*n + 4
	var maxk = 100
	var L = 900
	var dx = float64(L) / float64(n)
	x := make([][]float64, n)
	for i := range x {
		x[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		x[i][i] = 0
		for j := 0; j < n; j++ {
			if i != j {
				x[i][j] = float64(i+j) * dx
			}
		}
	}
	var T = 3
	dt := float64(T) / float64(maxk)
	alpha := D * dt / (dx * dx)
	FP := make([][]float64, nn)
	for i := range FP {
		FP[i] = make([]float64, nn)
	}

	F := make([][1]float64, nn)

	sus := make([][1]float64, n+2) // define initial populations
	for i := 0; i < n+2; i++ {
		sus[i][0] = 1 * sus0
	}
	for i := 1; i <= 3; i++ {
		sus[i][0] = 2
	}

	var susp = sus
	inf := make([][1]float64, n+2)
	for i := 0; i < n+2; i++ {
		inf[i][0] = float64(1 * inf0)
	}
	for i := 1; i <= 3; i++ {
		inf[i][0] = 48
	}
	var infp = inf

	time := make([][]float64, maxk)
	for i := range time {
		time[i] = make([]float64, maxk)
	}
	sustime := make([][]float64, maxk)
	for i := range sustime {
		sustime[i] = make([]float64, n+1)
	}
	inftime := make([][]float64, maxk)
	for i := range inftime {
		inftime[i] = make([]float64, n+1)
	}

	for k := 0; k < maxk; k++ {
		var aux []float64
		for _, arr := range susp {
			for _, item := range arr {
				aux = append(aux, item)
			}
		}
		for _, arr := range infp {
			for _, item := range arr {
				aux = append(aux, item)
			}
		}

		u := mat.NewDense(2*(n+2), 1, aux)
		m := 1
		var errors = 0.0
		for m = 0; m < 20; m++ {
			for i := 0; i < nn; i++ {
				if i >= 1 && i <= n {
					F[i][0] = sus[i][0] - susp[i][0] + dt*a*sus[i][0]*inf[i][0]
					FP[i][i] = 1 + dt*a*inf[i][0]
					FP[i][i+n+1] = dt * a * sus[i][0]
				}
				if i == n+2 {
					F[i][0] = inf[1][0] - infp[1][0] + b*dt*inf[1][0] - alpha*2*(-inf[1][0]+inf[2][0]) - a*dt*sus[1][0]*inf[1][0]
					FP[i][i] = 1 + b*dt + alpha*2 - a*dt*sus[1][0]
					FP[i][i+1] = -2 * alpha
					FP[i][1] = -a * dt * inf[1][0]

				}
				if i > n+2 && i < nn {
					i_shift := i - (n + 2)
					F[i][0] = inf[i_shift][0] - infp[i_shift][0] + b*dt*inf[i_shift][0] - alpha*(inf[i_shift-1][0]-2*inf[i_shift][0]+inf[i_shift][0]) - a*dt*sus[i_shift][0]*inf[i_shift][0]
					FP[i][i] = 1 + b*dt + alpha*2 - a*dt*sus[i_shift][0]
					FP[i][i-1] = -alpha
					FP[i][i] = -alpha
					FP[i][i_shift] = -a * dt * inf[i_shift][0]
				}
				if i == nn {
					F[i][0] = inf[n+1][0] - infp[n+1][0] + b*dt*inf[n+1][0] - alpha*2*(-inf[n+1][0]+inf[n][0]) - a*dt*sus[n+1][0]*inf[n+1][0]
					FP[i][i] = 1 + b*dt + alpha*2 - a*dt*sus[n+1][0]
					FP[i][i-1] = -2 * alpha
					FP[i][n+1] = -a * dt * inf[n+1][0]
				}
			}

			var result []float64
			for _, arr := range FP {
				for _, item := range arr {
					result = append(result, item)
				}
			}

			z := mat.NewDense(nn, nn, result)
			var du mat.Dense
			_ = du.Inverse(z)

			//if err != nil {
			//	log.Fatalf("z is not invertible: %v", err)
			//}

			u.Sub(u, &du)
			for i := 1; true; {
				sus[i][0] = u.At(i, 0)
			}
			var j = n + 2
			for i := 1; true; {
				if j < nn {
					inf[i][0] = u.At(j, 0)
					j = j + 1
				}
			}

			var result1 []float64
			for _, arr := range F {
				for _, item := range arr {
					result = append(result1, item)
				}
			}
			l := mat.NewDense(nn, 1, result)
			errors = mat.Norm(l, 2)
			if errors < 0.0001 {
				break
			}

		}

		time[k][k] = float64(k) * dt
		fmt.Println(time[k][k])
		fmt.Println(m)
		fmt.Println(errors)
		susp = sus
		infp = inf
		for i := range sustime {
			sustime[i][k] = sus[i][k]
		}
		for i := range inftime {
			inftime[i][k] = inf[i][k]
		}

	}

}
