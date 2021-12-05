package main

import (
	sonar "github.com/bmalcherek/AdventOfCode/01-sonar"
	steering "github.com/bmalcherek/AdventOfCode/02-steering"
	diagnostic "github.com/bmalcherek/AdventOfCode/03-diagnostic"
)

func main() {
	sonar.Run()
	steering.Run()
	diagnostic.Run()
}
