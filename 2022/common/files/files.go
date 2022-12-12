package files

import (
	op "aoc/common/operators"
	"bufio"
	"log"
	"os"
	"strings"
)

// Read and parse a the (demo|input).txt file
// Returns the lines of the file
func Read(filename string) []string {
	return GetLines(GetFile("../" + filename))
}

func ReadAndSplit(filename string, pattern string) [][]string {
	res := make([][]string, 0)
	lines := Read(filename)
	for _, l := range lines {
		res = append(res, strings.Split(l, pattern))
	}

	return res
}

func ReadAsMatrix(filename string) [][]int {
	lines := Read(filename)
	n := len(lines[0])
	m := make([][]int, n)

	for i, values := range lines {
		for j, v := range values {
			m[i] = append(m[i], op.Atoi(string(v)))
			j++
		}
	}

	return m
}

func GetFile(filename string) (file *os.File) {
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
