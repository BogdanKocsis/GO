package main

//This code is for susceptible/infected population.
//The infected may disperse in 1D via Fick's law.
//Newton's method is used.
//The full Jacobian matrix is defined.
//The linear steps are solved by A\d.

import (
	"fmt"
	//"gonum.org/v1/gonum/mat"
)

func main() {

	var sus0 = 60.0
	//var inf0 = 0
	a := 20.0 / 50.0
	fmt.Println(a)
	b := 1.0
	fmt.Println(b)
	var D = 1000.0
	fmt.Println(D)
	var n = 200
	var nn = 2*n + 2
	var maxk = 100
	var L = 900
	var dx = float64(L) / float64(n)
	fmt.Println(dx)
	x := make([]float64, n+1)

	for i := 0; i < n+1; i++ {
		x[i] = float64(i) * dx
	}

	fmt.Println(x)
	var T = 3
	dt := float64(T) / float64(maxk)
	alpha := D * dt / (dx * dx)
	fmt.Println(alpha)
	FP := make([][]float64, nn)
	for i := range FP {
		FP[i] = make([]float64, nn)
	}
	fmt.Println(len(FP))

	F := make([][1]float64, nn)
	fmt.Println(len(F))

	sus := make([][1]float64, n+1) // define initial populations
	for i := 0; i < n+1; i++ {
		sus[i][0] = 1 * sus0
	}
	fmt.Println(len(sus))

	for i := 0; i < 3; i++ {
		sus[i][0] = 2
	}

	var susp = sus
	fmt.Println(len(susp))

	inf := make([][1]float64, n+1)

	for i := 0; i < 3; i++ {
		inf[i][0] = 48
	}
	fmt.Println(len(inf))

	var infp = inf

	time := make([][]float64, maxk)
	for i := range time {
		time[i] = make([]float64, maxk)
	}
	fmt.Println(len(time))

	sustime := make([][]float64, n+1)
	for i := range sustime {
		sustime[i] = make([]float64, maxk)
	}
	fmt.Println(len(sustime))

	inftime := make([][]float64, n+1)
	for i := range inftime {
		inftime[i] = make([]float64, maxk)
	}
	fmt.Println(len(inftime))
	u := make([][1]float64, 2*(n+1))
	for i := 0; i < n+1; i++ {
		u[i][0] = susp[i][0]
	}
	var j = 0
	fmt.Println("u:", len(u))
	for i := n + 1; i < 2*(n+1); i++ {
		u[i][0] = infp[j][0]
		j = j + 1
	}

	for k := 0; k < maxk; k++ {
		//var aux []float64
		//for _, arr := range susp {
		//	for _, item := range arr {
		//		aux = append(aux, item)
		//	}
		//}
		//for _, arr := range infp {
		//	for _, item := range arr {
		//		aux = append(aux, item)
		//	}
		//}

		//	u := mat.NewDense(2*(n+2), 1, aux)
		u := make([][1]float64, 2*(n+1))
		for i := 0; i < n+1; i++ {
			u[i][0] = susp[i][0]
		}
		var j = 0
		fmt.Println("u:", len(u))
		for i := n + 1; i < 2*(n+1); i++ {
			u[i][0] = infp[j][0]
			j = j + 1
		}
		m := 1
		//var errors = 0.0
		for m = 0; m < 20; m++ {
			for i := 0; i < nn; i++ {
				if i >= 0 && i < n {
					F[i][0] = sus[i][0] - susp[i][0] + dt*a*sus[i][0]*inf[i][0]
					FP[i][i] = 1 + dt*a*inf[i][0]
					FP[i][i+n+1] = dt * a * sus[i][0]
				}
				if i == n+1 {
					F[i][0] = inf[1][0] - infp[1][0] + b*dt*inf[1][0] - alpha*2*(-inf[1][0]+inf[2][0]) - a*dt*sus[1][0]*inf[1][0]
					FP[i][i] = 1 + b*dt + alpha*2 - a*dt*sus[1][0]
					FP[i][i+1] = -2 * alpha
					FP[i][1] = -a * dt * inf[1][0]

				}
				if i > n+1 && i < nn {
					i_shift := i - (n + 1)
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

			//
			//		var result []float64
			//		for _, arr := range FP {
			//			for _, item := range arr {
			//				result = append(result, item)
			//			}
			//		}
			//
			//		z := mat.NewDense(nn, nn, result)
			//		var du mat.Dense
			//		_ = du.Inverse(z)
			//
			//		//if err != nil {
			//		//	log.Fatalf("z is not invertible: %v", err)
			//		//}
			//
			//		u.Sub(u, &du)
			//		for i := 1; true; {
			//			sus[i][0] = u.At(i, 0)
			//		}
			//		var j = n + 2
			//		for i := 1; true; {
			//			if j < nn {
			//				inf[i][0] = u.At(j, 0)
			//				j = j + 1
			//			}
			//		}
			//
			//		var result1 []float64
			//		for _, arr := range F {
			//			for _, item := range arr {
			//				result = append(result1, item)
			//			}
			//		}
			//		l := mat.NewDense(nn, 1, result)
			//		errors = mat.Norm(l, 2)
			//		if errors < 0.0001 {
			//			break
			//		}
			//
			//	}
			//
			//	time[k][k] = float64(k) * dt
			//	fmt.Println(time[k][k])
			//	fmt.Println(m)
			//	fmt.Println(errors)
			//	susp = sus
			//	infp = inf
			//	for i := range sustime {
			//		sustime[i][k] = sus[i][k]
			//	}
			//	for i := range inftime {
			//		inftime[i][k] = inf[i][k]
			//	}
			//
			//}
		}

	}
	fmt.Println(FP)
}
