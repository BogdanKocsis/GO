package main

//This code is for susceptible/infected population.
//The infected may disperse in 1D via Fick's law.
//Newton's method is used.
//The full Jacobian matrix is defined.
//The linear steps are solved by A\d.

func main() {

	var sus0 = 60.0
	var inf0 = 0
	var a = 20 / 50
	var b = 1
	var D = 1000
	var n = 200
	var nn = 2*n + 2
	var maxk = 100
	var L = 900
	var dx = L / n
	var x []int
	for i := 0; i < n; i++ {
		x[i] = dx * i
	}
	var T = 3
	var dt = T / maxk
	var alpha = D * dt / (dx * dx)
	FP := make([][]int, nn)
	for i := range FP {
		FP[i] = make([]int, nn)
	}

	F := make([][1]int, nn)
	sus := make([][1]int, n+1) // define initial populations
	for i := 0; i < n+1; i++ {
		sus[i][0] = int(1 * sus0)
	}
	for i := 1; i <= 3; i++ {
		sus[i][0] = 2
	}

	var susp = sus
	inf := make([][1]int, n+1)
	for i := 0; i < n+1; i++ {
		inf[i][0] = 1 * inf0
	}
	for i := 1; i <= 3; i++ {
		inf[i][0] = 48
	}
	var infp = inf

	time := make([][]int, maxk)
	for i := range time {
		time[i] = make([]int, maxk)
	}
	sustime := make([][]int, maxk)
	for i := range sustime {
		sustime[i] = make([]int, n+1)
	}
	inftime := make([][]int, maxk)
	for i := range inftime {
		inftime[i] = make([]int, n+1)
	}

	for k := 1; k <= maxk; k++ {
		//u := [susp; infp]
		for m := 1; m <= 20; m++ {
			for i := 1; i <= nn; i++ {
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
					i_shift := i - (n + 1)
					F[i][0] = inf[i_shift][0] - infp[i_shift][0] + b*dt*inf[i_shift][0] - alpha*(inf[i_shift-1][0]-2*inf[i_shift][0]+inf[i_shift+1][0]) - a*dt*sus[i_shift][0]*inf[i_shift][0]
					FP[i][i] = 1 + b*dt + alpha*2 - a*dt*sus[i_shift][0]
					FP[i][i-1] = -alpha
					FP[i][i+1] = -alpha
					FP[i][i_shift] = -a * dt * inf[i_shift][0]
				}
				if i == nn {
					F[i][0] = inf[n+1][0] - infp[n+1][0] + b*dt*inf[n+1][0] - alpha*2*(-inf[n+1][0]+inf[n][0]) - a*dt*sus[n+1][0]*inf[n+1][0]
					FP[i][i] = 1 + b*dt + alpha*2 - a*dt*sus[n+1][0]
					FP[i][i-1] = -2 * alpha
					FP[i][n+1] = -a * dt * inf[n+1][0]
				}
			}
		}

	}

}
