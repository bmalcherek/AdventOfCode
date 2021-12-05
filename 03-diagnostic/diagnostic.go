package diagnostic

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bmalcherek/AdventOfCode/utils"
)

func simpleDiagnostic(input []string) {
	onesCount := make([]int, len(input[0]))
	for _, reading := range input {
		for i := 0; i < len(reading); i++ {
			if string(reading[i]) == "1" {
				onesCount[i] += 1
			}
		}
	}
	gammaStr := ""
	epsilonStr := ""
	for _, count := range onesCount {
		if count > len(input)/2 {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 32)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 32)
	fmt.Printf("05. Gamma: %d, Epsilon: %d, Power reading: %d\n", gamma, epsilon, gamma*epsilon)
}

func Run() {
	file, err := os.Open("./03-diagnostic/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := utils.ReadFileToStringArray(file)

	simpleDiagnostic(input)
}
