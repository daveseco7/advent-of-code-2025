package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			log.Fatalf("something went wrong closing the file: %v", closeErr)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadLinesAsInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			log.Fatalf("something went wrong closing the file: %v", closeErr)
		}
	}(file)

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, strConvErr := strconv.Atoi(scanner.Text())
		if strConvErr != nil {
			return nil, strConvErr
		}
		lines = append(lines, n)
	}
	return lines, scanner.Err()
}
