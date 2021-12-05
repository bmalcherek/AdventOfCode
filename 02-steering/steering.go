package steering

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/bmalcherek/AdventOfCode/utils"
)

type simplePosition struct {
	x int
	y int
}

func simpleSteering(input []string) {
	upRegex := regexp.MustCompile(`^up`)
	downRegex := regexp.MustCompile(`^down`)
	forwardRegex := regexp.MustCompile(`^forward`)
	numRegex := regexp.MustCompile(`[0-9]+`)

	currentPosition := simplePosition{
		x: 0,
		y: 0,
	}
	for _, command := range input {
		change, err := strconv.ParseInt(numRegex.FindAllString(command, 1)[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case upRegex.MatchString(command):
			currentPosition.y -= int(change)
		case downRegex.MatchString(command):
			currentPosition.y += int(change)
		case forwardRegex.MatchString(command):
			currentPosition.x += int(change)
		}
	}

	fmt.Printf("03. Current submarine position: depth: %d, x: %d, task result: %d\n", currentPosition.y, currentPosition.x, currentPosition.x*currentPosition.y)
}

type advancedPosition struct {
	x   int
	y   int
	aim int
}

func advancedSteering(input []string) {
	upRegex := regexp.MustCompile(`^up`)
	downRegex := regexp.MustCompile(`^down`)
	forwardRegex := regexp.MustCompile(`^forward`)
	numRegex := regexp.MustCompile(`[0-9]+`)

	currentPosition := advancedPosition{
		x:   0,
		y:   0,
		aim: 0,
	}
	for _, command := range input {
		change, err := strconv.ParseInt(numRegex.FindAllString(command, 1)[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case upRegex.MatchString(command):
			currentPosition.aim -= int(change)
		case downRegex.MatchString(command):
			currentPosition.aim += int(change)
		case forwardRegex.MatchString(command):
			currentPosition.x += int(change)
			currentPosition.y += int(change) * currentPosition.aim
		}
	}

	fmt.Printf("04. Current submarine position: depth: %d, x: %d, aim: %d, task result: %d\n", currentPosition.y, currentPosition.x, currentPosition.aim, currentPosition.x*currentPosition.y)

}

func Run() {
	file, err := os.Open("./02-steering/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := utils.ReadFileToStringArray(file)

	simpleSteering(input)
	advancedSteering(input)
}
