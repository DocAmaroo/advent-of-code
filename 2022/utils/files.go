package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) (file *os.File) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func GetLines(file *os.File) []string {
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
	return GetLines(ReadFile("../" + filename))
}

func RemoveBlankSpaces(l []string) [][]string {
	res := make([][]string, len(l))
	for i, v := range l {
		res[i] = strings.Split(v, " ")
	}

	return res
}
