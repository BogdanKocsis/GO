package main

import (
	fmt "fmt"
	"math"
	"strings"
	//"math/big"
	//"strconv"
	//"strings"
)

type Rational struct {
	Numa, Numi int
}

func (r Rational) getNuma() int {
	return r.Numa
}

func (r Rational) getNumi() int {
	return r.Numi
}

func (r Rational) newInstance(x, y int) Rational {
	t := Rational{x, y}
	return t
}

func (r Rational) add(t Rational) Rational {

	r.Numa = (r.Numa * t.Numi) + (t.Numa * r.Numi)
	r.Numi = r.Numi * t.Numi

	return r.simplify()
}

func (r Rational) subtract(t Rational) Rational {

	r.Numa = (r.Numa * t.Numi) - (t.Numa * r.Numi)
	r.Numi = r.Numi * t.Numi

	return r.simplify()

}

func (r Rational) multiply(t Rational) Rational {

	r.Numa = r.Numa * t.Numa
	r.Numi = r.Numi * t.Numi
	return r.simplify()
}

func (r Rational) multiplyInt(t int) Rational {

	r.Numa = r.Numa * t

	return r.simplify()
}

func (r Rational) divideInt(t int) Rational {

	aux := Rational{1, t}
	return r.multiply(aux)
}

func (r Rational) subtractInt(t int) Rational {
	aux := Rational{t, 1}
	return r.subtract(aux)
}

func (r Rational) addInt(t int) Rational {

	aux := Rational{t, 1}
	return r.add(aux)
}

func (r Rational) addNuma(t int) Rational {

	return r.addInt(t)
}

func (r Rational) addNumi(t int) Rational {

	aux := Rational{1, t}
	return r.add(aux)
}

func (r Rational) addNumaAndNumi(t int) Rational {

	aux := Rational{t, t}
	return r.add(aux)
}

func (r Rational) subtractNuma(t int) Rational {

	return r.subtractInt(t)
}

func (r Rational) subtractNumi(t int) Rational {

	aux := Rational{1, t}
	return r.subtract(aux)
}

func (r Rational) subtractNumaAndNumi(t int) Rational {

	aux := Rational{t, t}
	return r.subtract(aux)
}

func (r Rational) isNull() bool {
	return r.Numa == 0
}

func (r Rational) getRealValue() float32 {

	var result = float32(r.Numa) / float32(r.Numi)

	return result
}

func (r Rational) getAbsValue() Rational {

	if r.Numi > 0 {
		return r
	} else {
		r.Numi = -r.Numi
		return r
	}

}

func (r Rational) divide(t Rational) Rational {
	aux := Rational{t.Numi, t.Numa}

	return r.multiply(aux)
}

func (r Rational) simplify() Rational {
	c := cmmdc(r.Numa, r.Numi)
	r.Numa /= c
	r.Numi /= c
	return r
}

func (r Rational) pow(n int) Rational {

	r.Numa = int(math.Pow(float64(r.Numa), float64(n)))
	r.Numi = int(math.Pow(float64(r.Numi), float64(n)))

	return r
}

func (r Rational) biggerThan(t Rational) bool {

	if r.getRealValue() > t.getRealValue() {
		return true
	} else {
		return false
	}

}

func (r Rational) smallerThan(t Rational) bool {
	if r.biggerThan(t) == false {
		return true
	} else {
		return false
	}

}

func (r Rational) equals(t Rational) bool {
	a := r.simplify()
	b := t.simplify()

	return a.Numi == b.Numi &&
		a.Numa == b.Numa
}

func (r Rational) inverse() Rational {

	return Rational{r.Numi, r.Numa}
}

func (r Rational) isNatural() bool {

	return r.Numi == 1
}

// 0.4324 => 4324 / 1000 => simplificare
func getFromFloat32(x float32) Rational {

	i := fmt.Sprint(x)
	if strings.Contains(i, ".") {
		zeros := len(strings.Split(fmt.Sprintf("%v", x), ".")[1])
		numi := math.Pow(10, float64(zeros))
		numa := x * float32(numi)
		result := Rational{int(numa), int(numi)}
		return result.simplify()

	} else {
		return Rational{int(x), 1}
	}

	//
	//dec := 0
	//for float64(x) != math.Trunc(float64(x)) {
	//	dec++
	//	x = x * 10
	//}
	//return Rational{int(x), int(math.Pow10(dec))}

	//s := fmt.Sprintf("%f", x)
	//r, _ := new(big.Rat).SetString(s)
	//
	//var s1 = r.String()
	//s2 := strings.Split(s1, "/")
	//
	//if a, err := strconv.ParseInt(s2[0], 10, 64); err == nil {
	//	result.Numa = int(a)
	//}
	//if b, err := strconv.ParseInt(s2[1], 10, 64); err == nil {
	//	result.Numi = int(b)
	//}
}

func getNegative(r Rational) Rational {

	r.Numi = r.Numi * -1
	return r
}

func getSquareRoot(r Rational) float64 {

	result := r.getRealValue()

	return math.Sqrt(float64(result))
}

func cmmdc(a, b int) int {
	for a%b != 0 {
		r := a % b
		a = b
		b = r
	}
	return b
}

func main() {

	//se considera ca numitorul este diferit de 0
	number1 := Rational{1, 3}

	number2 := Rational{2, 5}

	fmt.Println("Add: ", number1.add(number2))
	fmt.Println("Subtract: ", number1.subtract(number2))
	fmt.Println("Multiply: ", number1.multiply(number2))
	fmt.Println("Multiply Int: ", number1.multiplyInt(7))
	fmt.Println("Divide Int: ", number1.divideInt(7))
	fmt.Println("Subtract Int: ", number1.subtractInt(7))
	fmt.Println("Add Int: ", number1.addInt(7))
	fmt.Println("Add Numa: ", number1.addNuma(7))
	fmt.Println("Add Numi: ", number1.addNumi(7))
	fmt.Println("Add Numa & Numi ", number1.addNumaAndNumi(7))
	fmt.Println("Subtract Numa: ", number1.subtractNuma(7))
	fmt.Println("Subtract Numi: ", number1.subtractNumi(7))
	fmt.Println("Subtract Numa and Numi: ", number1.subtractNumaAndNumi(7))
	fmt.Println("Is Null: ", number1.isNull())
	fmt.Println("Get Real Value: ", number1.getRealValue())
	fmt.Println("Get Abs Value: ", number1.getAbsValue())
	fmt.Println("Divide: ", number1.divide(number2))
	fmt.Println("Pow: ", number1.pow(3))
	fmt.Println("Bigger than : ", number1.biggerThan(number2))
	fmt.Println("Smaller Than: ", number1.smallerThan(number2))
	fmt.Println("Equals: ", number1.equals(number2))
	fmt.Println("Is Natural: ", number1.isNatural())
	fmt.Println("Inverse: ", number1.inverse())
	fmt.Println("Get From Float:", getFromFloat32(-2.25))
	fmt.Println("Get Negative :", getNegative(number2))
	fmt.Println("Get Square Root :", getSquareRoot(number2))
}
