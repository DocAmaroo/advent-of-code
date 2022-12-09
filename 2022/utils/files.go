package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filename string) (file *os.File) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func ParseByLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadLines(filename string) []string {
	return ParseByLines(ReadFile("../" + filename))
}
