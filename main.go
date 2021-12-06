package main

import (
	sonar "github.com/bmalcherek/AdventOfCode/01-sonar"
	steering "github.com/bmalcherek/AdventOfCode/02-steering"
	diagnostic "github.com/bmalcherek/AdventOfCode/03-diagnostic"
	bingo "github.com/bmalcherek/AdventOfCode/04-bingo"
)

func main() {
	sonar.Run()
	steering.Run()
	diagnostic.Run()
	bingo.Run()
}
