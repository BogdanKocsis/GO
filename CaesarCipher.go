package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandomValueInRange() uint8 {

	return uint8(rand.Int()%25 + 1)
}

func read() string {
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	return value
}

func caesarEncrypt(text string) {
	cipher := ""

	keyFile, err := os.Create("key.txt")
	if err != nil {
		fmt.Print(err)
	}

	cryptFile, err := os.Create("crypt.txt")
	if err != nil {
		fmt.Print(err)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(text)-1; i++ {
		if text[i] >= 'A' && text[i] <= 'Z' {
			shift := getRandomValueInRange()
			fmt.Fprintln(keyFile, shift)

			if text[i]+shift > 'Z' {
				cipher += string(text[i] + shift + 'A' - 'Z' - 1)
				fmt.Fprint(cryptFile, string(text[i]+shift+'A'-'Z'-1))
			} else {
				cipher += string(text[i] + shift)
				fmt.Fprint(cryptFile, string(text[i]+shift))
			}
		} else {
			cipher += " "
			fmt.Fprint(cryptFile, " ")
		}
	}
	fmt.Println(cipher)
}

func decrypt() {

	message := ""

	keyFile, err := os.Open("key.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//keyFile.Close()

	var shifts []string
	scanner := bufio.NewScanner(keyFile)
	for scanner.Scan() {
		shifts = append(shifts, scanner.Text())
	}
	fmt.Println(shifts)

	cryptFile, err := ioutil.ReadFile("crypt.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	cipher := string(cryptFile)
	fmt.Println(cipher)
	k := 0
	for i := 0; i < len(cipher); i++ {
		if cipher[i] >= 'A' && cipher[i] <= 'Z' {
			j, _ := strconv.ParseUint(shifts[k], 10, 32)
			shift := uint8(j)
			if cipher[i]-shift < 'A' {
				shift = 26 - shift
				fmt.Println(shift)
				message += string(cipher[i] + shift)
			} else {
				message += string(cipher[i] - shift)
			}
			k++
		} else {
			message += " "
		}
	}
	fmt.Println(message)

}

func main() {

	//str := read()
	//caesarEncrypt(str)
	decrypt()
}
