package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./day3/input")

	lines := strings.Split(string(data), "\n")

	strLen := len(lines[0])
	var oxygenIndices []int
	var co2Indices []int

	for i := 0; i < len(lines); i++ {
		oxygenIndices = append(oxygenIndices, i)
		co2Indices = append(co2Indices, i)
	}

	for strPos := 0; strPos < strLen; strPos++ {
		if len(oxygenIndices) > 1 {
			onesOxygenIndices, zeroesOxygenIndices := partition(oxygenIndices, lines, strPos)
			if len(onesOxygenIndices) >= len(zeroesOxygenIndices) {
				oxygenIndices = onesOxygenIndices
			} else {
				oxygenIndices = zeroesOxygenIndices
			}
		}

		if len(co2Indices) > 1 {
			onesCo2Indices, zeroesCo2Indices := partition(co2Indices, lines, strPos)
			if len(zeroesCo2Indices) <= len(onesCo2Indices) {
				co2Indices = zeroesCo2Indices
			} else {
				co2Indices = onesCo2Indices
			}
		}
	}

	oxygen, _ := strconv.ParseInt(lines[oxygenIndices[0]], 2, 64)
	co2, _ := strconv.ParseInt(lines[co2Indices[0]], 2, 64)

	println(lines[oxygenIndices[0]], lines[co2Indices[0]])
	println(oxygen, co2)
	println(oxygen * co2)
}

func partition(indices []int, lines []string, strPos int) ([]int, []int) {
	var onesIndices []int
	var zerosIndices []int
	for _, index := range indices {
		lineAtIndex := lines[index]
		symbol := []rune(lineAtIndex)[strPos]
		if symbol == '1' {
			onesIndices = append(onesIndices, index)
		} else {
			zerosIndices = append(zerosIndices, index)
		}
	}
	return onesIndices, zerosIndices
}
