package main

import (
	"fmt"
)

func isValidSudoko(sudoko [9][9]uint) bool {
	a := [3][9]uint{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			pos := sudoko[i][j]
			line := a[0][i]
			column := a[1][j]
			square := a[2][(3*(i/3) + j/3)]
			var bitPos uint = 1 << pos
			if (line&bitPos) == bitPos || (column&bitPos) == bitPos || (square&bitPos) == bitPos {
				return false
			}
			a[0][i] = line | bitPos
			a[1][j] = column | bitPos
			a[2][3*(i/3)+j/3] = square | bitPos
		}
	}

	return true
}

func main() {

	var sudoko = [9][9]uint{{7, 9, 2, 1, 5, 4, 3, 8, 6},
		{6, 4, 3, 8, 2, 7, 1, 5, 9},
		{8, 5, 1, 3, 9, 6, 7, 2, 4},
		{2, 6, 5, 9, 7, 3, 8, 4, 1},
		{4, 8, 9, 5, 6, 1, 2, 7, 3},
		{3, 1, 7, 4, 8, 2, 9, 6, 5},
		{1, 3, 6, 7, 4, 8, 5, 9, 2},
		{9, 7, 4, 2, 1, 5, 6, 3, 8},
		{5, 2, 8, 6, 3, 9, 4, 1, 7}}

	var sudokoMistake = [9][9]uint{{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5}}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", sudoko[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Is Valid? ", isValidSudoko(sudoko))

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", sudokoMistake[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Is Valid? ", isValidSudoko(sudokoMistake))
}
