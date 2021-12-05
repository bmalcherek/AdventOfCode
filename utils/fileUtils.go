package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileToIntArray(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	result := []int{}
	for scanner.Scan() {
		val, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, int(val))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func ReadFileToStringArray(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
