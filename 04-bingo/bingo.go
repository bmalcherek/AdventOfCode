package bingo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type bingoBoard struct {
	board [][]string
}

func Run() {
	file, err := os.Open("./04-bingo/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	bingoNumbers := scanner.Text()

	bingoBoards := []bingoBoard{}
	bingoBoard := bingoBoard{
		board: [][]string{},
	}
	readingBoard := false
	numberLineRegex := regexp.MustCompile(`^(\s\d|\d{2})`)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 && !readingBoard {
			readingBoard = true
			splittedLine := strings.Split(scanner.Text(), `\s`)
			bingoBoard.board = append(bingoBoard.board, splittedLine)
		} else if numberLineRegex.MatchString(scanner.Text()) {
			splittedLine := strings.Split(scanner.Text(), `\s`)
			for i := range splittedLine {
				fmt.Println(splittedLine[i], i)
			}
			fmt.Println(splittedLine, len(splittedLine))
		}
	}

	fmt.Println(bingoNumbers, bingoBoards, len(bingoBoard.board))
}
