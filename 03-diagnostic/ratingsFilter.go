package diagnostic

import (
	"regexp"
)

func oxygenGeneratorRatingFilter(input []string) string {
	filteredInputs := input
	inputRegexStr := "^"
	var oxygenGeneratorRatingStr string
	for i := 0; i < len(input[0]); i++ {
		onesCount := countOnesOccurences(filteredInputs)
		if float64(onesCount[i]) >= float64(len(filteredInputs))/2 {
			inputRegexStr += "1"
		} else {
			inputRegexStr += "0"
		}
		inputRegex := regexp.MustCompile(inputRegexStr)
		matchedInputs := []string{}
		for j := range input {
			if inputRegex.MatchString(input[j]) {
				matchedInputs = append(matchedInputs, input[j])
			}
		}

		if len(matchedInputs) == 1 {
			oxygenGeneratorRatingStr = matchedInputs[0]
			break
		}
		filteredInputs = matchedInputs
	}

	return oxygenGeneratorRatingStr
}

func co2ScrubberRatingFilter(input []string) string {
	filteredInputs := input
	inputRegexStr := "^"
	var co2ScrubberRatingStr string
	for i := 0; i < len(input[0]); i++ {
		onesCount := countOnesOccurences(filteredInputs)
		if float64(onesCount[i]) >= float64(len(filteredInputs))/2 {
			inputRegexStr += "0"
		} else {
			inputRegexStr += "1"
		}
		inputRegex := regexp.MustCompile(inputRegexStr)
		matchedInputs := []string{}
		for j := range input {
			if inputRegex.MatchString(input[j]) {
				matchedInputs = append(matchedInputs, input[j])
			}
		}

		if len(matchedInputs) == 1 {
			co2ScrubberRatingStr = matchedInputs[0]
			break
		}
		filteredInputs = matchedInputs
	}

	return co2ScrubberRatingStr
}
