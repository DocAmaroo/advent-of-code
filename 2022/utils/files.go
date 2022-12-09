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

// Read and parse a the (demo|input).txt file
// Returns the lines of the file
func Read(filename string) []string {
	return ParseByLines(ReadFile("../" + filename))
}
