package main

import "fmt"

func weekDay(number int8) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid value"
	}
}

func weekDayTwo(number int8) string {
	switch {
	case number == 1:
		return "Sunday"
	case number == 2:
		return "Monday"
	case number == 3:
		return "Tuesday"
	case number == 4:
		return "Wednesday"
	case number == 5:
		return "Thursday"
	case number == 6:
		return "Friday"
	case number == 7:
		return "Saturday"
	default:
		return "Invalid value"
	}
}

func test(number int8) string {
	var value string

	switch {
	case number == 1:
		value = "1"
		fallthrough
	default:
		value = "Well, your code drop here anyway"
		return value
	}

}

func main() {
	day := weekDay(3)
	fmt.Println(day)

	dayTwo := weekDay(3)
	fmt.Println(dayTwo)

	test := test(1)
	fmt.Println(test)
}