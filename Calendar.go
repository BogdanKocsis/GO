package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

const layoutRo = "02-01-2006"

type Date struct {
	year  int
	month int
	day   int
}

//SetYear
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) ReadDate() {

	var year, month, day int
	fmt.Print("Year:\t\t")
	readValue(&year)
	_ = d.SetYear(year)
	fmt.Print("Month:\t\t")
	readValue(&month)
	_ = d.SetMonth(month)
	fmt.Print("Day:\t\t")
	readValue(&day)
	_ = d.SetDay(day)

	fmt.Println("Date: ", d.Day(), "-", d.Month(), "-", d.Year())

}

func ToDate(year, month, day int) time.Time { // convert input to time.Date format
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func isLeap(y int) bool {

	if y%100 != 0 && y%4 == 0 || y%400 == 0 {
		return true
	}
	return false
}

func countLeapYears(date time.Time) (leaps int) {

	// returns year, month,
	// date of a time object
	y, m, _ := date.Date()

	if m <= 2 {
		y--
	}
	leaps = y/4 + y/400 - y/100
	return leaps
}

func getDifference(a, b time.Time) (days int) {

	// month-wise days
	monthDays := [12]int{31, 28, 31, 30, 31,
		30, 31, 31, 30, 31, 30, 31}

	// extracting years, months,
	// days of two dates
	y1, m1, d1 := a.Date()
	y2, m2, d2 := b.Date()

	// totalDays since the
	// beginning = year*365 + number_of_days
	totalDays1 := y1*365 + d1

	// adding days of the months
	// before the current month
	for i := 0; i < (int)(m1)-1; i++ {
		totalDays1 += monthDays[i]
	}

	// counting leap years since
	// beginning to the year "a"
	// and adding that many extra
	// days to the totaldays
	totalDays1 += countLeapYears(a)

	// Similar procedure for second date
	totalDays2 := y2*365 + d2

	for i := 0; i < (int)(m2)-1; i++ {
		totalDays2 += monthDays[i]
	}

	totalDays2 += countLeapYears(b)

	// Number of days between two days
	days = totalDays2 - totalDays1

	return days

}

func offsetDays(a time.Time) int {

	y, m, d := a.Date()

	offset := d
	fmt.Println(offset)
	if m-1 == 11 {
		offset += 335
	}
	if m-1 == 10 {
		offset += 304
	}

	if m-1 == 9 {
		offset += 273
	}
	if m-1 == 8 {
		offset += 243
	}
	if m-1 == 7 {
		offset += 212
	}
	if m-1 == 6 {
		offset += 181
	}
	if m-1 == 5 {
		offset += 151
	}
	if m-1 == 4 {
		offset += 120
	}
	if m-1 == 3 {
		offset += 90
	}
	if m-1 == 2 {
		offset += 59
	}
	if m-1 == 1 {
		offset += 31
	}
	fmt.Println(offset)

	if isLeap(y) && m > 2 {
		offset += 1
	}
	fmt.Println(offset)
	return offset
}

func revoffsetDays(offset int, y int) (d int, m int) {

	month := [13]int{0, 31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31}

	if isLeap(y) {
		month[2] = 29
	}
	var i int
	for i = 1; i <= 12; i++ {
		if offset <= month[i] {
			break
		}
		offset = offset - month[i]

	}

	d = offset
	m = i
	return d, m
}

func addDays(a time.Time, days int) {

	y, _, _ := a.Date()
	offset1 := offsetDays(a)
	fmt.Println(offset1)
	remDays := 0

	if isLeap(y) {
		remDays = 366 - offset1
	} else {
		remDays = 365 - offset1
	}

	var y1 int
	var offset2 int

	if days <= remDays {
		y1 = y
		offset2 = offset1 + days
	} else {
		days -= remDays
		y1 = y + 1
		var y1days int
		if isLeap(y1) {
			y1days = 366
		} else {
			y1days = 365
		}
		for days >= y1days {
			days -= y1days
			y1++
			if isLeap(y1) {
				y1days = 366
			} else {
				y1days = 365
			}
		}
		offset2 = days
	}
	day, month := revoffsetDays(offset2, y1)
	fmt.Println(day, month, y1)
}

func timeSubDays(t1, t2 time.Time) int {

	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		return -1
	}
	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := t1y == t2y && t1m == t2m && t1d == t2d

		if isSameDay {

			return 0
		} else {
			return 1
		}

	} else { // equal or more than 24

		if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times
			return int(hours / 24)
		} else { // more than 24 hours
			return int(hours/24) + 1
		}
	}
}

func timeSubYearsMonthsWeeksDays(days float64) {

	years := math.Floor(days / 365) // approximates down
	months := math.Floor((days - years*365) / 30)
	weeks := math.Floor((days - years*365 - months*30) / 7)
	day := days - years*365 - months*30 - weeks*7 //calculates the number of days left in the calculation

	fmt.Printf("The diffence between of these 2 days: %v years, %v months, %v weeks, %v days. \n", years, months, weeks, day)
	fmt.Println()
}

//func AddDays(t2 time.Time) {
//
//	layout := layoutRo
//	var numberOfDays int
//	fmt.Print("Number of days:\t\t")
//	readValue(&numberOfDays)
//	date := t2.AddDate(0, 0, numberOfDays)
//	fmt.Printf("%v \n", date.Format(layout))
//	fmt.Println()
//}

func SubDays(t2 time.Time) {

	layout := layoutRo
	var numberOfDays int
	fmt.Print("Number of days:\t\t")
	readValue(&numberOfDays)
	date := t2.AddDate(0, 0, numberOfDays*-1)
	fmt.Printf("%v \n", date.Format(layout))
	fmt.Println()
}

func readValue(value *int) {

	_, err := fmt.Scanf("%d", value)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	fmt.Println("Write first date")
	date1 := Date{}
	date1.ReadDate()
	//print(date1.Year(),"-",date1.Month(),"-",date1.Day())

	fmt.Println("Write second data")
	date2 := Date{}
	date2.ReadDate()
	//print(date2.Year(),"-",date2.Month(),"-",date2.Day())

	t1 := ToDate(date1.Year(), date1.Month(), date1.Day())
	t2 := ToDate(date2.Year(), date2.Month(), date2.Day())

	fmt.Println(" 1. Number of days  between two dates \n 2. Difference of two dates in years, months, weeks and day \n 3. Add number of days on a date \n 4. Subtract number of days on a date \n ")

	var caseNumber int
	fmt.Print("Case Number:\t\t")
	var _, _ = fmt.Scanln(&caseNumber)

	switch caseNumber {
	case 1:

		if t1.After(t2) {
			t1, t2 = t2, t1
		}
		days := getDifference(t1, t2)
		fmt.Println(days)

	case 2:
		days := timeSubDays(t1, t2)
		timeSubYearsMonthsWeeksDays(float64(days))

	case 3:
		//AddDays(t2)
		addDays(t1, 366)
	case 4:
		SubDays(t2)
	}

}
