package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./day1_2/source")

	lines := strings.Split(string(data), "\n")
	ints := make([]int64, len(lines))
	for i, line := range lines {
		ints[i], _ = strconv.ParseInt(line, 0, 64)
	}

	counter := 0
	var previousSum int64

	for i := 2; i < len(ints); i++ {
		currentSum := ints[i] + ints[i-1] + ints[i-2]
		println(currentSum)
		if previousSum != 0 && currentSum > previousSum {
			counter++
		}
		previousSum = currentSum
	}
	println(counter)
}
