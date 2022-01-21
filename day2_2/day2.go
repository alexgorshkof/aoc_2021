package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./day2_2/input")

	lines := strings.Split(string(data), "\n")
	var horizontalPos, verticalPos, aim int64
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		dir := lineSplit[0]
		amount, _ := strconv.ParseInt(lineSplit[1], 0, 64)
		switch dir {
		case "down":
			aim += amount
		case "up":
			aim -= amount
		case "forward":
			horizontalPos += amount
			verticalPos += amount * aim
		}
	}
	println(horizontalPos * verticalPos)
}
