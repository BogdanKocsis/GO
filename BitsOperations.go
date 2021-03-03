package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func isPowerOfTwo(n int) bool {
	return (n > 0) && ((n & (n - 1)) == 0)
}

func swap(a *int, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
	return
}

func encodeDecode(data string) {

	var mask int64 = 63
	stringMask := strconv.FormatInt(mask, 2)
	aux := stringMask

	divide := len(data) / len(stringMask)
	remainder := len(data) % len(stringMask)

	for i := 0; i < divide; i++ {
		stringMask += aux
	}
	if remainder != 0 {
		stringMask += stringMask[0:remainder]
	}
	fmt.Println("Mask: ", stringMask)

	information, err := strconv.ParseInt(data, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	code, err := strconv.ParseInt(stringMask, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	x := information ^ code
	fmt.Println("Encoded information: ", strconv.FormatInt(x, 2))
	y := x ^ code
	fmt.Println("Decoded information: ", strconv.FormatInt(y, 2))

}

func bitExtracted(number int, p int, t int) int {

	//number = (number >> t-1) & (1 << p-1)
	number = (number >> (t - 1)) & (1<<p - 1)
	return number
}

func convertToHex(number int64) {

	var s = ""
	for number > 0 {
		residue := number % (1 << 4)
		s += strconv.FormatInt(residue, 16)
		number = number / (1 << 4)
	}
	fmt.Println(s)
}

func main() {

	fmt.Printf("Menu\n 1.Is Power of Two\n 2.Bit Swap\n 3.Encode/Decode \n 4.Extracted Number\n 5.Convert to Hex \n")
	var i int
	fmt.Printf("Write the option: ")
	_, _ = fmt.Scanf("%d", &i)
	switch i {
	case 1:
		fmt.Printf("Write the number: ")
		var i int
		_, _ = fmt.Scanf("%d", &i)
		if isPowerOfTwo(i) {
			fmt.Println("Number", i, "is power of two")
		} else {
			fmt.Println("Number", i, "is not power of two")
		}

	case 2:
		var x = 5
		var y = 7
		fmt.Println("Before: ", x, y)
		swap(&x, &y)
		fmt.Print("After: ", x, y)

	case 3:
		data, err := ioutil.ReadFile("test.txt")
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fmt.Println("Contents of file:", string(data))
		encodeDecode(string(data))

	case 4:

		fmt.Println("The extracted number is: ", bitExtracted(171, 5, 2))

	case 5:
		fmt.Printf("Write the number to convert: ")
		var i int64
		_, _ = fmt.Scanf("%d", &i)
		convertToHex(i)

	}

}
