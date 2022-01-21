package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./day1/source")

	lines := strings.Split(string(data), "\n")
	counter := 0
	var previous int64

	for index, line := range lines {
		if index == 0 {
			continue
		}
		current, _ := strconv.ParseInt(line, 0, 16)
		if current > previous {
			counter++
		}

		previous = current
	}
	println(counter)
}
