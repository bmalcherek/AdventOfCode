package sonar

import (
	"fmt"
	"log"
	"os"

	"github.com/bmalcherek/AdventOfCode/utils"
)

func simpleSonar(input []int) {
	depthIncrease := 0
	lastNum := input[0]
	for _, num := range input {
		if num > lastNum {
			depthIncrease += 1
		}
		lastNum = num
	}
	fmt.Printf("01. Number of depth increases: %d\n", depthIncrease)
}

func slidingSumSonar(input []int) {
	depthIncrease := 0
	lastSum := input[0] + input[1] + input[2]
	for i := 0; i < len(input)-2; i++ {
		sum := input[i] + input[i+1] + input[i+2]
		if sum > lastSum {
			depthIncrease += 1
		}
		lastSum = sum
	}
	fmt.Printf("02. Number of depth increases: %d\n", depthIncrease)

}

func Run() {
	file, err := os.Open("./01-sonar/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := utils.ReadFileToIntArray(file)

	simpleSonar(input)
	slidingSumSonar(input)
}
