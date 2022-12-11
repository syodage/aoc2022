package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var Debug bool

func Printlnf(format string, a ...any) {
	if Debug {
		Printf(fmt.Sprintf("%s\n", format), a...)
	}
}
func Printf(format string, a ...any) {
	if Debug {
		fmt.Printf(fmt.Sprintf(format, a...))
	}
}

func ReadLines(path string) []string {
	return readLines(path)
}

func ReadLinesWithIgnoreHead(path string, ignoreHead int) []string {
	res := readLines(path)
	return res[ignoreHead:]
}

func ReadHead(path string, n int) []string {
	res := readLines(path)
	return res[:n]
}

func readLines(path string) []string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var res []string
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		res = append(res, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
