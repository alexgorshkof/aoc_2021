package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./day2/input")

	lines := strings.Split(string(data), "\n")
	var horizontalPos, verticalPos int64
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		dir := lineSplit[0]
		amount, _ := strconv.ParseInt(lineSplit[1], 0, 64)
		switch dir {
		case "down":
			verticalPos += amount
		case "up":
			verticalPos -= amount
		case "forward":
			horizontalPos += amount
		}
	}
	println(horizontalPos * verticalPos)
}
