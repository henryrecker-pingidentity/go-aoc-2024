package day1

import (
	"errors"
	"strconv"
)

func Part1(input []string) {
	values := []int{}
	for _, line := range input {
		firstDigit := ""
		lastDigit := ""
		for _, digit := range line {
			strDigit := string(digit)
			_, err := strconv.Atoi(strDigit)
			if err == nil {
				if firstDigit == "" {
					firstDigit = strDigit
				}
				lastDigit = strDigit
			}
		}
		val, _ := strconv.Atoi(firstDigit + lastDigit)
		values = append(values, val)
	}

	result := 0
	for _, val := range values {
		result += val
	}
	println("Result:", result)
}

func getDigit(runes []rune, startIndex int) (string, error) {
	if startIndex < len(runes)-2 {
		strToCheck := string([]rune{runes[startIndex], runes[startIndex+1], runes[startIndex+2]})
		switch strToCheck {
		case "one":
			return "1", nil
		case "two":
			return "2", nil
		case "six":
			return "6", nil
		}
	}
	if startIndex < len(runes)-3 {
		strToCheck := string([]rune{runes[startIndex], runes[startIndex+1], runes[startIndex+2], runes[startIndex+3]})
		switch strToCheck {
		case "four":
			return "4", nil
		case "five":
			return "5", nil
		case "nine":
			return "9", nil
		}
	}
	if startIndex < len(runes)-4 {
		strToCheck := string([]rune{runes[startIndex], runes[startIndex+1], runes[startIndex+2], runes[startIndex+3], runes[startIndex+4]})
		switch strToCheck {
		case "three":
			return "3", nil
		case "seven":
			return "7", nil
		case "eight":
			return "8", nil
		}
	}
	return "", errors.New("No valid digit value found")
}

func Part2(input []string) {
	values := []int{}
	for _, line := range input {
		runes := []rune(line)
		firstDigit := ""
		lastDigit := ""
		for i, r := range runes {
			// Try parsing as int
			strVal := string(r)
			_, err := strconv.Atoi(strVal)
			if err == nil {
				if firstDigit == "" {
					firstDigit = strVal
				}
				lastDigit = strVal
			} else {
				// Check if this is one of the written-out digits
				digit, err := getDigit(runes, i)
				if err == nil {
					if firstDigit == "" {
						firstDigit = digit
					}
					lastDigit = digit
				}
			}
		}
		val, _ := strconv.Atoi(firstDigit + lastDigit)
		values = append(values, val)
	}

	result := 0
	for _, val := range values {
		result += val
	}
	println("Result:", result)
}
