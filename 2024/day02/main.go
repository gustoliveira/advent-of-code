package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputFile = "input.txt"

func main() {
	args := os.Args[1:]

	if len(args) == 2 && args[1] == "base" {
		inputFile = "input.base.txt"
	}

	if args[0] == "part1" {
		part1()
		return
	}

	if args[0] == "part2" {
		part2()
		return
	}
}

func part1() {
	lines, err := readFile()
	if err != nil {
		fmt.Println("Error reading file")
	}

	safeLevelsCount := 0

	for _, line := range lines {
		levels := strings.Split(line, " ")

		if len(levels) <= 1 {
			safeLevelsCount++
			continue
		}

		safeLevel := isSafeLevel(levels)

		if safeLevel {
			safeLevelsCount++
		}
	}

	fmt.Println(safeLevelsCount)
}

func isSafeLevel(levels []string) bool {
	safeLevel := true

	isIncreasing := toInt(levels[0]) > toInt(levels[1])

	for i := 1; i < len(levels); i++ {
		currentLevel := toInt(levels[i])
		previousLevel := toInt(levels[i-1])

		if currentLevel == previousLevel {
			safeLevel = false
			break
		}

		levelDiff := absolute(currentLevel - previousLevel)

		if levelDiff < 1 || levelDiff > 3 {
			safeLevel = false
			break
		}

		if isIncreasing && (currentLevel > previousLevel) {
			safeLevel = false
			break
		}

		if !isIncreasing && (currentLevel < previousLevel) {
			safeLevel = false
			break
		}
	}

	return safeLevel
}

func part2() {
	lines, err := readFile()
	if err != nil {
		fmt.Println("Error reading file")
	}

	safeLevelsCount := 0

	for _, line := range lines {
		levels := strings.Split(line, " ")

		if len(levels) <= 1 {
			safeLevelsCount++
			continue
		}

		safeLevel := isSafeLevel(levels)

		if safeLevel {
			safeLevelsCount++
			continue
		}

		safeLevel = false

		for index := range levels {
			removedLevels := removeElementAt(levels, index)

			safeLevelNew := isSafeLevel(removedLevels)

			if safeLevelNew {
				safeLevel = true
				break
			}
		}

		if safeLevel {
			safeLevelsCount++
		}
	}

	fmt.Println(safeLevelsCount)
}

func readFile() ([]string, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}

func toInt(n string) int {
	value, _ := strconv.Atoi(n)
	return value
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func removeElementAt(s []string, i int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:i]...)
	return append(ret, s[i+1:]...)
}
