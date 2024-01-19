package main

import (
	"adventofcode2024/solutions/day1"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world.")

	input, err := readFileLines("input/day1/part1.txt")
	if err != nil {
		println("Error opening input file: ", err.Error())
		return
	}

	day1.Part1(input)
}

// Thanks https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
func readFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
