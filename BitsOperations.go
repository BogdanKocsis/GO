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

func bitFlow(data string) {

	//output, err := strconv.ParseInt(string(data), 2, 64)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	var mask int64 = 63
	stringMask := strconv.FormatInt(mask, 2)
	fmt.Println("Mask: ", strconv.FormatInt(mask, 2))
	//x := output ^ mask
	//fmt.Println("Info with mask: ",strconv.FormatInt(x, 2))
	var step int = 0
	var x string = ""
	for i := 0; i < len(data); i = i + len(stringMask) {
		s := data[i : len(stringMask)+step]
		output, _ := strconv.ParseInt(string(s), 2, 64)
		aux := output ^ mask
		x = x + strconv.FormatInt(aux, 2)
		step += len(stringMask)
		fmt.Printf("%s\n", s)
		if len(stringMask)+step > len(data) {
			break
		}

	}
	//if len(data) % len(stringMask) != 0 {
	//	fmt.Println("ok")
	//	var lastCharacters = data[len(data)-(len(data)%len(stringMask)):]
	//
	//	output, _ := strconv.ParseInt(string(lastCharacters), 2, 64)
	//	stringMask = stringMask[0:(len(data)%len(stringMask))]
	//	mask, _ := strconv.ParseInt(string(stringMask), 2, 64)
	//	aux := output ^ mask
	//	println(output)
	//	x = x + strconv.FormatInt(aux, 2)
	//}

	fmt.Println("Mask: ", x, len(x), len(data))

	//var t int64
	//t = x ^ mask
	//fmt.Println("Same Info:", strconv.FormatInt(t, 2))

}

func bitExtracted(number int, p int, t int) int {

	//number = (number >> t-1) & (1 << p-1)
	number = (number >> (t - 1)) & (1<<p - 1)
	return number
}

func convertToHex(number int64) {

	var s string = ""
	for number > 0 {
		residue := number % (1 << 4)
		s += strconv.FormatInt(residue, 16)
		number = number / (1 << 4)
	}
	fmt.Println(s)
}

func main() {

	fmt.Printf("Menu\n 1.Is Power of Two\n 2.Bit Swap\n 3. \n 4.Extracted Number\n 5.Convert to Hex \n")
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
		bitFlow(string(data))

	case 4:

		fmt.Println("The extracted number is: ", bitExtracted(171, 5, 2))

	case 5:
		fmt.Printf("Write the number to convert: ")
		var i int64
		_, _ = fmt.Scanf("%d", &i)
		convertToHex(i)

	}

}
