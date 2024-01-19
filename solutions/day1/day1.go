package day1

import (
	"strconv"
)

func Part1(input []string) {
	values := []int{}
	for _, line := range input {
		firstDigit := '_'
		lastDigit := '_'
		for _, digit := range line {
			_, err := strconv.Atoi(string(digit))
			if err == nil {
				if firstDigit == '_' {
					firstDigit = digit
				}
				lastDigit = digit
			}
		}
		val, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
		values = append(values, val)
	}

	result := 0
	for _, val := range values {
		result += val
	}
	println("Result:", result)
}
