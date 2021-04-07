package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	return value
}

func initializeMatrix() [5][5]string {

	alph := "ABCDEFGHIJKLMNOPRSTUVWXYZ"
	letters := [5][5]string{}

	k := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			letters[i][j] = string(alph[k])
			k++
		}
	}

	return letters
}

func playfairEncrypt(data [5][5]string, letters string) string {

	l1x := 0
	l2x := 0
	l1y := 0
	l2y := 0
	found := false
	l1 := string([]rune(letters)[0])
	l2 := string([]rune(letters)[1])
	if !strings.Contains(l1, " ") && !strings.Contains(l2, " ") {
		for l1x < 5 {
			l1y = 0
			for l1y < 5 {
				if l1 == data[l1x][l1y] {
					found = true
					break
				}
				l1y++
			}
			if found == true {
				found = false
				break
			}
			l1x++
		}

		for l2x < 5 {
			l2y = 0
			for l2y < 5 {
				if l2 == data[l2x][l2y] {
					found = true
					break
				}
				l2y++
			}
			if found == true {
				break
			}
			l2x++
		}

		//case 1
		if l2x == l1x {
			l2x++
			l1x++
			if l2x > 4 {
				l2x = 0
			}
			if l1x > 4 {
				l1x = 0
			}
		}
		//case 2
		if l2y == l1y {
			l2y++
			l1y++
			if l1y > 4 {
				l1y = 0
			}
			if l2y > 4 {
				l2y = 0
			}
		}
		//case 3
		if l2y != l1y && l2x != l1x {
			holder := l1x
			l1x = l2x
			l2x = holder
		}

		returnVal := data[l1x][l1y] + data[l2x][l2y]
		return returnVal
	} else {
		return letters
	}
}

func splitLetters(data [5][5]string, keyword string) string {

	cipher := ""

	length := float64(len(keyword))
	lastCharacter := keyword[len(keyword)-1:]
	if math.Mod(length, 2) != 0 {
		r, size := utf8.DecodeLastRuneInString(keyword)
		if r == utf8.RuneError && (size == 0 || size == 1) {
			size = 0
		}
		keyword = keyword[:len(keyword)-size]
	}

	for i := 0; i < len(keyword)-1; i = i + 2 {
		cipher += playfairEncrypt(data, keyword[i:i+2])
	}
	if math.Mod(length, 2) != 0 {
		cipher += lastCharacter
	}
	return cipher
}

func main() {

	data := initializeMatrix()
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("Enter the message you want to encrypt")
	str := readString()
	str = strings.ToUpper(str)
	cipher := splitLetters(data, str[:len(str)-1])
	fmt.Println(cipher)
}
