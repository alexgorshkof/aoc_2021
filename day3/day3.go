package main

import (
	"math"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./day3/input")

	lines := strings.Split(string(data), "\n")

	strLen := len(lines[0])
	var gammaRate, epsilonRate int
	for i := 0; i < strLen; i++ {
		onesCounter := 0
		for _, line := range lines {
			symbol := []rune(line)[i]
			if symbol == '1' {
				onesCounter++
			}
		}

		currentPow := int(math.Pow(2, float64(strLen-i-1)))
		if onesCounter > len(lines)/2 {
			gammaRate += currentPow
		} else {
			epsilonRate += currentPow
		}
		println(currentPow, gammaRate, epsilonRate)
	}
	println(gammaRate * epsilonRate)
}
